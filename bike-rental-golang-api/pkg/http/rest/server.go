package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
)

type DomainServices struct {
	ListService ports.ListService
}

type httpServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Server struct {
	httpServer httpServer
}

func NewServer(host string, port uint16, validator ports.Validator, dServices DomainServices) *Server {
	r := chi.NewRouter()

	r.Mount("/api/v1", NewV1Router(validator, dServices))

	return &Server{httpServer: &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}}
}

func (s *Server) Serve(context context.Context) error {
	var err error

	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		errChan <- s.httpServer.ListenAndServe()
	}()

	for {
		select {
		case err = <-errChan:
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}

			return nil
		case <-context.Done():
			return context.Err()
		default:
		}
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("shutting down http server gracefully")

	return s.httpServer.Shutdown(ctx)
}
