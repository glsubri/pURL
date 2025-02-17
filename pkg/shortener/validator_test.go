package shortener

import "testing"

func TestUrlsValidity(t *testing.T) {
	tcs := []struct {
		url         string
		expectedErr bool
	}{
		{url: "https://www.google.com", expectedErr: false},
		{url: "http://www.google.com", expectedErr: false},
		{url: "ftp://www.google.com", expectedErr: false},
		{url: "https://www.google.com/this/is/a/test", expectedErr: false},
		{url: "http://google.com", expectedErr: false},
		{url: "notaurl", expectedErr: true},
	}

	for _, tc := range tcs {
		err := validateURL(tc.url)
		if err != nil && !tc.expectedErr {
			t.Fatalf("unexpected error has occured: %s\n", err)
		}
		if err == nil && tc.expectedErr {
			t.Fatalf("error should have occured but didn't for url: '%s'.\n", tc.url)
		}
	}
}
