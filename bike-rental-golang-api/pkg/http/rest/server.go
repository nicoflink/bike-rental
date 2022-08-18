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
	RentService ports.RentService
}

type httpServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Server struct {
	httpServer httpServer
}

// NewServer creates a new server for REST API.
// It defines the version of the API that are provided to the customer.
func NewServer(host string, port uint16, validator ports.Validator, dServices DomainServices) (*Server, error) {
	if host == "" || port == 0 || validator == nil {
		return nil, errors.New("NewServer dependencies are not fulfilled")
	}

	r := chi.NewRouter()

	v1Router, err := NewV1Router(validator, dServices)
	if err != nil {
		return nil, fmt.Errorf("rest.NewServer: %w", err)
	}

	r.Mount("/api/v1", v1Router)

	return &Server{httpServer: &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}}, nil
}

// Serve starts serving the REST API.
// It calls ListenAndServe of the underlying HTTP server and listen to the provided context for any interruption.
// If an error occurs other than http.ErrServerClosed then the error will be returned.
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

// Shutdown shuts down the underlying http server.
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("shutting down http server gracefully")

	return s.httpServer.Shutdown(ctx)
}
