package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// Prefix of the authorization header for an anonymous user
const ANONYMOUS_PREFIX = "Anonymous"

// Prefix of the authorization header for an authenticated user
const BEARER_PREFIX = "Bearer"

// Key for the "uid" stored in context
const UID = "uid"

// Authentication Middleware
// Calls the next handler with `uid` in the `context` which can be used as a unique
// id for any user.
// Returns `http.StatusUnauthorized` to the client if the authorization token is not
// found or is invalid.
func AuthMiddleware(next http.Handler) http.Handler {
	tokenValidator := GetJWTValidator()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization_header := r.Header.Get("Authorization")

		// Default error message
		errorMessage := "Invalid authorization"

		if strings.HasPrefix(authorization_header, ANONYMOUS_PREFIX) {
			// For anonymous authentication

			// Get the token from the header
			uid := strings.Split(authorization_header, ANONYMOUS_PREFIX)[1]
			uid = strings.Trim(uid, " \n\r")

			// If a valid token is present
			if len(uid) > 0 {
				// Add given token as uid to the context
				ctx := context.WithValue(r.Context(), UID, uid)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			errorMessage = "Anonymous token not found"
		} else if strings.HasPrefix(authorization_header, BEARER_PREFIX) {
			// For bearer authentication

			// Get the JWT token from the header
			token := strings.Split(authorization_header, BEARER_PREFIX)[1]
			token = strings.Trim(token, " \n\r")

			// if a token is present
			if len(token) > 0 {
				// Validate the JWT token and get claims
				tokenDecoded, err := tokenValidator.ValidateToken(context.TODO(), token)

				if err == nil {
					tokenDecoded := tokenDecoded.(*validator.ValidatedClaims)

					// Use the "sub" claim as the uid
					uid := tokenDecoded.RegisteredClaims.Subject

					// if a valid uid is present
					if len(uid) > 0 {
						// Add the given "sub" as the uid to the context
						ctx := context.WithValue(r.Context(), UID, uid)
						next.ServeHTTP(w, r.WithContext(ctx))
						return
					}
				} else {
					log.Println(err)
				}
			}
			errorMessage = "Invalid authentication token"
		}

		// If the next handler wasn't called, authentication failed
		// return Unauthorized http error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, errorMessage)))
	})
}

// Returns a JWT Token validator
func GetJWTValidator() *validator.Validator {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithAllowedClockSkew(time.Minute),
	)

	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	return jwtValidator
}
