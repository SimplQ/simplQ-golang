package api

type ValidationError struct {
    Field   string
    Message string
}

type Validator interface {
    Validate() (ValidationError, bool)
}
