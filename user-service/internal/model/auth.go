package model

import "log"

const (
	CustomerRole = "customer"
	AdminRole    = "admin"

	// Admin credentials
	AdminEmail    = "admi@foodstore.com"
	AdminPassword = "sequence0"
)

func init() {
	log.Printf("Admin credentials loaded - Email: %s", AdminEmail)
}

// AuthorizationError represents an error related to authorization
type AuthorizationError struct {
	Message string
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

var (
	ErrUnauthorized = &AuthorizationError{Message: "unauthorized"}
	ErrForbidden    = &AuthorizationError{Message: "forbidden"}
	ErrNotFound     = &AuthorizationError{Message: "not found"}
)
