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
		shortURL := ""
		for !saved {
			shortURL = shorten(url, s.host, length)
			saved = s.save(url, shortURL)
		}

		return shortURL, nil
	}
}

func (s *InMemory) save(rawURL string, shortURL string) bool {
	if _, exists := s.urls[shortURL]; exists {
		return false
	}

	s.urls[shortURL] = rawURL
	return true
}

func (s *InMemory) OriginalURL(ctx context.Context, shortURL string) (string, error) {
	originalURL, exists := s.urls[shortURL]
	if !exists {
		return "", ErrShortURLDoesNotExist
	}

	return originalURL, nil
}
