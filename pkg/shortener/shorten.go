package shortener

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func shorten(URL string, host string, length int) string {
	randPath := generateRandomString(length)

	h := host
	if host[len(host)-1] != '/' {
		h = host + "/"
	}

	return h + randPath
}

func generateRandomString(length int) string {
	seeded := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		randIndex := seeded.Intn(len(charset))
		sb.WriteByte(charset[randIndex])
	}

	return sb.String()
}
