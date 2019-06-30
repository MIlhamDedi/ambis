package constants

import "errors"

var (
	ErrInternalError = errors.New("Internal Error")

	// Authentication
	ErrExpiredToken = errors.New("Token Expired")
)
