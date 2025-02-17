package shortener

import "testing"

func TestShortenReturnsCorrectLength(t *testing.T) {
	host := "https://www.subri.ch/"
	tcs := []struct {
		url string
		len int
	}{
		{"https://www.website.com", 10},
		{"https://www.google.com/this/is/a/link", 4},
		{"https://www.subri.ch/this/is/a/very/very/very/long/link", 4},
	}

	for _, tc := range tcs {
		shortened := shorten(tc.url, host, tc.len)
		shortLen := len(shortened) - len(host)

		if shortLen != tc.len {
			t.Fatalf("url '%s' was shortened to '%s', expected length was %d but got %d",
				tc.url,
				shortened,
				tc.len,
				shortLen,
			)
		}
	}
}
