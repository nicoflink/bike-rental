package bikes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/render"
	"github.com/nicoflink/bike-rental/pkg/list"
)

type Service interface {
	GetAllBikes(ctx context.Context) ([]list.Bike, error)
}

type handler struct {
	service   Service
	validator ports.Validator
}

func NewBikesRouter(listService Service) chi.Router {
	r := chi.NewRouter()

	h := handler{
		service: listService,
	}

	r.Get("/", h.listBikes)

	return r
}

func (h handler) listBikes(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	bikes, err := h.service.GetAllBikes(ctx)
	if err != nil {
		return
	}

	render.Json(writer, mapToBikesJsonResponse(bikes))
}
