package admin

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Server struct {
	mux       *http.ServeMux
	assetFS   fs.FS
	templates *template.Template
}

func NewServer(opts ...Option) (*Server, error) {
	s := &Server{
		mux: http.NewServeMux(),
	}
	for _, o := range opts {
		o(s)
	}
	return s, nil
}

// ServeHandler implments http.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) init() error {
	if s.assetFS == nil {
		return errors.New("asset filesystem is required")
	}
	s.mux.Handle("/css/", http.FileServer(http.FS(s.assetFS)))
	s.mux.Handle("/img/", http.FileServer(http.FS(s.assetFS)))
	s.mux.HandleFunc("/api/status", s.handleAPIStatus)
	s.mux.HandleFunc("/", s.handleRoot)

	var err error
	s.templates, err = template.New("").ParseFS(s.assetFS, "*.html")
	if err != nil {
		return fmt.Errorf("parsing admin templates: %w", err)
	}

	return nil
}

func (s *Server) handleAPIStatus(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	data := struct{}{}
	if err := s.templates.Lookup("index.html").Execute(w, data); err != nil {
		log.Warn().Err(err).Msg("executing / template")
	}
}
