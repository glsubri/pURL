package shortener

import (
	"context"
	"testing"
)

func TestInMemorySaveRetrieveMatches(t *testing.T) {
	inmem := NewInMemory("localhost:3000")

	originalURL := "https://www.subri.ch/this/is/a/link"
	shortCode, err := inmem.Shorten(context.Background(), originalURL, 10)
	if err != nil {
		t.Fatalf("should not have received error: %s\n", err)
	}

	savedOriginalURL, err := inmem.OriginalURL(context.Background(), shortCode)
	if err != nil {
		t.Fatalf("should not have received error: %s\n", err)
	}

	if originalURL != savedOriginalURL {
		t.Fatalf("saved original URL '%s' should be '%s'\n", savedOriginalURL, originalURL)
	}
}
