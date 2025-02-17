package shortener

import "net/url"

// validateURL checks if a given raw URL string is valid.
// It returns nil if the URL is valid, or an appropriate error if it's invalid.
func validateURL(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ErrInvalidURL
	}

	if u.Host == "" {
		return ErrMissingHost
	}

	return nil
}
