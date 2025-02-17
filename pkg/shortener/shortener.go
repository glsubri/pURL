package shortener

import "context"

// Shortener is the interface for URL shortening and forwarding operations.
// It provides methods to both shorten a URL and retrieve the original URL from a shortened one.
type Shortener interface {
	// Shorten creates a shortened URL from an original URL.
	// It returns an error if the URL is invalid or cannot be shortened.
	Shorten(ctx context.Context, url string, length int) (string, error)

	// OriginalURL returns the original URL for a given shortened URL.
	// It returns an error if the shortened URL is not found or invalid.
	OriginalURL(ctx context.Context, shortURL string) (string, error)
}
