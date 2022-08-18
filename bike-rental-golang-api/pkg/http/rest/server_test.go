package rest

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/list"
	"github.com/nicoflink/bike-rental/pkg/rent"
)

type httpServerMock struct {
	mockListenAndServe func() error
	mockShutdown       func(ctx context.Context) error
}

func (h httpServerMock) ListenAndServe() error {
	return h.mockListenAndServe()
}

func (h httpServerMock) Shutdown(ctx context.Context) error {
	return h.mockShutdown(ctx)
}

type httpServerMockOption func(*httpServerMock)

func withListenAndServeError(err error) httpServerMockOption {
	return func(mock *httpServerMock) {
		mock.mockListenAndServe = func() error {
			return err
		}
	}
}

func withListenAndServeSleep(d time.Duration) httpServerMockOption {
	return func(mock *httpServerMock) {
		mock.mockListenAndServe = func() error {
			time.Sleep(d)

			return nil
		}
	}
}

func withShutdownErr(err error) httpServerMockOption {
	return func(mock *httpServerMock) {
		mock.mockShutdown = func(ctx context.Context) error {
			return err
		}
	}
}

func newHttpServerMock(options ...httpServerMockOption) httpServerMock {
	mock := &httpServerMock{
		mockListenAndServe: func() error {
			return nil
		},
		mockShutdown: func(ctx context.Context) error {
			return nil
		},
	}

	for _, o := range options {
		o(mock)
	}

	return *mock
}

func TestNewServer(t *testing.T) {
	type args struct {
		host      string
		port      uint16
		validator ports.Validator
		dServices DomainServices
	}

	validDomainServices := DomainServices{
		ListService: list.Service{},
		RentService: rent.Service{},
	}

	tests := []struct {
		name    string
		input   args
		wantErr bool
	}{
		{
			name:    "Happy Case: NewServer",
			input:   args{host: "localhost", port: 8080, validator: validator.New(), dServices: validDomainServices},
			wantErr: false,
		},
		{
			name:    "Error: host is empty",
			input:   args{port: 8080, validator: validator.New(), dServices: validDomainServices},
			wantErr: true,
		},
		{
			name:    "Error: port is empty",
			input:   args{host: "localhost", validator: validator.New(), dServices: validDomainServices},
			wantErr: true,
		},
		{
			name:    "Error: validator is empty",
			input:   args{host: "localhost", port: 8080, validator: nil, dServices: validDomainServices},
			wantErr: true,
		},
		{
			name:    "Error: NewV1Router returns error",
			input:   args{host: "localhost", port: 8080, validator: validator.New(), dServices: DomainServices{}},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, err := NewServer(tc.input.host, tc.input.port, tc.input.validator, tc.input.dServices)
			if err != nil && !tc.wantErr {
				t.Errorf("Expected no error but got: %v", err)
			}

			if err == nil && tc.wantErr {
				t.Error("Expected error but nil")
			}

			if err == nil && server == nil {
				t.Errorf("Expected server to be initialized but got nil")
			}
		})
	}
}

func TestServer_Serve(t *testing.T) {
	type contextType uint8

	const (
		contextTODO contextType = iota
		contextCancel
		contextTimeout
	)

	tests := []struct {
		name       string
		fieldInput httpServer
		ctxType    contextType
		wantErr    bool
	}{
		{
			name:       "Happy Case: ListenAndServes returns http.ErrServerClosed",
			fieldInput: newHttpServerMock(withListenAndServeError(http.ErrServerClosed)),
			ctxType:    contextTODO,
			wantErr:    false,
		},
		{
			name:       "Error: ListenAndServes return error",
			fieldInput: newHttpServerMock(withListenAndServeError(errors.New("TestError"))),
			ctxType:    contextTODO,
			wantErr:    true,
		},
		{
			name:       "Error: Context Cancel Err",
			fieldInput: newHttpServerMock(withListenAndServeSleep(60 * time.Second)),
			ctxType:    contextCancel,
			wantErr:    true,
		},
		{
			name:       "Error: Context Timeout",
			fieldInput: newHttpServerMock(withListenAndServeSleep(60 * time.Second)),
			ctxType:    contextTimeout,
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			var ctx context.Context
			var cancel context.CancelFunc

			switch tc.ctxType {
			case contextTODO:
				ctx = context.TODO()
			case contextCancel:
				ctx, cancel = context.WithCancel(context.TODO())
				defer cancel()
			case contextTimeout:
				ctx, cancel = context.WithTimeout(context.TODO(), 0*time.Second)
				defer cancel()
			default:
				ctx = context.TODO()
			}

			s := Server{httpServer: tc.fieldInput}

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()

				err := s.Serve(ctx)
				if err != nil && !tc.wantErr {
					t.Errorf("Expected no error but got: %v", err)
				}

				if err == nil && tc.wantErr {
					t.Error("Expected error but nil")
				}
			}()

			if tc.ctxType == contextCancel {
				cancel()
			}

			wg.Wait()
		})
	}
}

func TestServer_Shutdown(t *testing.T) {
	tests := []struct {
		name       string
		fieldInput httpServer
		wantErr    bool
	}{
		{
			name:       "Happy Case: Shutdown",
			fieldInput: newHttpServerMock(),
			wantErr:    false,
		},
		{
			name:       "Error: httpServer returns error",
			fieldInput: newHttpServerMock(withShutdownErr(errors.New("ShutdownError"))),
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := Server{httpServer: tc.fieldInput}

			err := s.Shutdown(context.TODO())
			if err != nil && !tc.wantErr {
				t.Errorf("Expected no error but got: %v", err)
			}

			if err == nil && tc.wantErr {
				t.Error("Expected error but nil")
			}
		})
	}
}
