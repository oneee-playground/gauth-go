package gauth

import (
	"fmt"
	"net/http"
)

// GauthErr is returned when http request's status >= 400.
type GauthErr struct {
	Code    int
	Message string
}

func (e *GauthErr) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(e.Code), e.Message)
}

func newGauthErr(code int, message string) *GauthErr {
	return &GauthErr{code, message}
}
