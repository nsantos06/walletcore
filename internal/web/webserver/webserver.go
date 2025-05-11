package webserver

import (
	"net/http"
	"path"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router chi.Router
	Handlers map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(router chi.Router, handlers map[string]http.HandlerFunc, webServerPort string) *WebServer {
	return &WebServer{
		Router: router,
		Handlers: handlers,
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Post(path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}