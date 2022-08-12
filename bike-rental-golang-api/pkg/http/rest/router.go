package rest

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/bikes"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/rents"
)

func NewV1Router(validator ports.Validator, dServices DomainServices) chi.Router {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(middleware.UserCtx)

	r.Mount("/bikes", bikes.NewBikesRouter(dServices.ListService))
	r.Mount("/rents", rents.NewRentsRouter(dServices.RentService, validator))

	return r
}
