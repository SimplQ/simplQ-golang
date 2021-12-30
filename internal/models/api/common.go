package api

type ValidationError struct {
    Fields []string
    Message string
}

type Validator interface {
    Validate() (ValidationError, bool)
}
