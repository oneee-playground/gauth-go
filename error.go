package gauth

import (
	"fmt"
	"net/http"
)

// GauthErr is returned when http request's status >= 400.
type GauthErr struct {
	Code int
}

func (e *GauthErr) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, http.StatusText(e.Code))
}

func newGauthErr(code int) *GauthErr {
	return &GauthErr{code}
}
