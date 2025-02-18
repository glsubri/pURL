package addhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	errorMsgUnsupportedMediaType = "media type is not supported"
	errorMsgMalformedData        = "malformed data"
	errorMsgShortenServiceError  = "could not shorten given url"
	errorMsgMissingOriginalURL   = "missing 'original_url' field"
)

type HTTPError struct {
	Status  int
	Message string
	Err     error
}

func (e *HTTPError) Error() string {
	if e.Err == nil {
		return e.Message
	}

	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func (e *HTTPError) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": e.Message,
	})
}

func ErrUnsupportedMediaType(err error) *HTTPError {
	return &HTTPError{
		http.StatusUnsupportedMediaType,
		errorMsgUnsupportedMediaType,
		err,
	}
}

func ErrInvalidInput(err error) *HTTPError {
	return &HTTPError{
		http.StatusBadRequest,
		errorMsgMalformedData,
		err,
	}
}

func ErrInvalidInputMissingOriginalURL(err error) *HTTPError {
	return &HTTPError{
		http.StatusBadRequest,
		errorMsgMissingOriginalURL,
		err,
	}
}

func ErrShortenService(err error) *HTTPError {
	return &HTTPError{
		http.StatusInternalServerError,
		errorMsgShortenServiceError,
		err,
	}
}
