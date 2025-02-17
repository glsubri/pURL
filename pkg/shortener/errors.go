package shortener

import "errors"

var (
	ErrInvalidURL           = errors.New("invalid URL format")
	ErrMissingHost          = errors.New("URL must have a host")
	ErrShortURLDoesNotExist = errors.New("given short URL does not exist")
)
