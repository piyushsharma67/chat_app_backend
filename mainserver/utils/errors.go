package utils

import "errors"

var (
	ErrErrorOccurred               = errors.New("error occurred during authentication")
	ErrorInvalidpassword           = errors.New("invalid password")
	ErrorUserAlreadyExists         = errors.New("user already exists")
	ErrUserForbidden               = errors.New("login required")
	ErrErrorOccurredAuthentication = errors.New("error occurred during authentication")
)
