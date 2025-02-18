package addhandler

import (
	"encoding/json"
	"net/http"

	"github.com/glsubri/pURL/pkg/shortener"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type Handler struct {
	shortener      shortener.Shortener
	shortURLLength int
}

func New(shortener shortener.Shortener, shortURLLength int) *Handler {
	return &Handler{
		shortener:      shortener,
		shortURLLength: shortURLLength,
	}
}

type AddRequest struct {
	OriginalURL string `json:"original_url" form:"original_url" schema:"original_url"`
}

type AddResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addReq, httpErr := parseData(r)
		if httpErr != nil {
			httpErr.Respond(w)
			return
		}

		if httpErr = addReq.Validate(); httpErr != nil {
			httpErr.Respond(w)
			return
		}

		shortURL, err := h.shortener.Shorten(r.Context(), addReq.OriginalURL, h.shortURLLength)
		if err != nil {
			ErrShortenService(err).Respond(w)
			return
		}

		resp := AddResponse{
			OriginalURL: addReq.OriginalURL,
			ShortURL:    shortURL,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func parseData(r *http.Request) (AddRequest, *HTTPError) {
	if len(r.URL.Query()) > 0 {
		return parseURLValues(r)
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		return parseJSON(r)
	case "application/x-www-form-urlencoded":
		if err := r.ParseForm(); err != nil {
			return AddRequest{}, ErrInvalidInput(err)
		}
		return parseURLValues(r)
	}

	return AddRequest{}, ErrUnsupportedMediaType(nil)
}

func parseJSON(r *http.Request) (AddRequest, *HTTPError) {
	var req AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return AddRequest{}, ErrInvalidInput(err)
	}

	return req, nil
}

func parseURLValues(r *http.Request) (AddRequest, *HTTPError) {
	var req AddRequest
	if err := decoder.Decode(&req, r.PostForm); err != nil {
		return AddRequest{}, ErrInvalidInput(err)
	}

	return req, nil
}

func (req *AddRequest) Validate() *HTTPError {
	if req.OriginalURL == "" {
		return ErrInvalidInputMissingOriginalURL(nil)
	}

	return nil
}
