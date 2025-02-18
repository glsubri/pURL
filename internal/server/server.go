package server

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/glsubri/pURL/internal/server/handlers/addhandler"
	"github.com/glsubri/pURL/pkg/shortener"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	shortURLLength int
	shortener      shortener.Shortener
	router         chi.Router
}

func NewServer(
	shortener shortener.Shortener,
	shortURLLength int,
) *Server {
	s := &Server{
		shortURLLength: shortURLLength,
		shortener:      shortener,
		router:         chi.NewRouter(),
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	// Middleware
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	addHandler := addhandler.New(s.shortener, s.shortURLLength)
	// Routes
	s.router.HandleFunc("/add", addHandler.Handle)

	// Catch-all
	s.router.HandleFunc("/*", s.forward)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) forward(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	orignalURL, err := s.shortener.OriginalURL(r.Context(), shortCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("shortener: %s", err), http.StatusBadRequest)
	}

	http.Redirect(w, r, orignalURL, http.StatusPermanentRedirect)
}
