package jewelry

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(port string, h http.Handler) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpserver.ListenAndServe()
}

func (s *Server) ShutDown(context context.Context) error {
	return s.httpserver.Shutdown(context)
}
