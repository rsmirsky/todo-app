package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func(s *Server) Run(post string) error {
	s.httpServer = &http.Server{
		Addr: ":" ,
		MaxHeaderBytes: 1<<20, //1mb
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}