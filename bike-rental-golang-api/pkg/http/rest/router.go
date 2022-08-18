package rest

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/bikes"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/rents"
)

// NewV1Router defines the routes for version 1 of the REST API.
func NewV1Router(validator ports.Validator, dServices DomainServices) (router chi.Router, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rest.NewV1Router: %w", err)
		}
	}()

	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(middleware.UserCtx)

	bikesRouter, err := bikes.NewBikesRouter(dServices.ListService, validator)
	if err != nil {
		return nil, err
	}

	rentsRouter, err := rents.NewRentsRouter(dServices.RentService, validator)
	if err != nil {
		return nil, err
	}

	r.Mount("/bikes", bikesRouter)
	r.Mount("/rents", rentsRouter)

	return r, nil
}
