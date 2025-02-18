package shortener

import (
	"context"
)

type InMemory struct {
	host string
	urls map[string]string
}

func NewInMemory(host string) *InMemory {
	return &InMemory{
		host: host,
		urls: make(map[string]string),
	}
}

func (s *InMemory) Shorten(ctx context.Context, url string, length int) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		err := validateURL(url)
		if err != nil {
			return "", err
		}

		// If url is valid, generate name and save it
		saved := false
		shortCode := ""
		for !saved {
			shortCode = generateRandomString(length)
			saved = s.save(url, shortCode)
		}

		return shortCode, nil
	}
}

func (s *InMemory) save(rawURL string, shortCode string) bool {
	if _, exists := s.urls[shortCode]; exists {
		return false
	}

	s.urls[shortCode] = rawURL
	return true
}

func (s *InMemory) OriginalURL(ctx context.Context, shortCode string) (string, error) {
	originalURL, exists := s.urls[shortCode]
	if !exists {
		return "", ErrShortURLDoesNotExist(shortCode)
	}

	return originalURL, nil
}
