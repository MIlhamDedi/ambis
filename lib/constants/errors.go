package constants

import "errors"

// Error messages
const (
	S_NO_ERROR       = "Success"
	S_INTERNAL_ERROR = "Internal Error"

	// Authentication
	S_EXPIRED_TOKEN  = "Token Expired"
	S_NO_MATCH_FOUND = "No Matching Credentials"
)

var (
	ErrInternalError = errors.New(S_INTERNAL_ERROR)
	ErrExpiredToken  = errors.New(S_EXPIRED_TOKEN)
)
