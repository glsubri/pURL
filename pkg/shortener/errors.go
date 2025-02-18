package shortener

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidURL  = errors.New("invalid URL format")
	ErrMissingHost = errors.New("URL must have a host")
)

func ErrShortURLDoesNotExist(shortURL string) error {
	return fmt.Errorf("short URL '%s' does not exist", shortURL)
}
