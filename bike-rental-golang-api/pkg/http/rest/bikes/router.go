package bikes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/render"
)

type handler struct {
	service   ports.ListService
	validator ports.Validator
}

func NewBikesRouter(listService ports.ListService) chi.Router {
	r := chi.NewRouter()

	h := handler{
		service: listService,
	}

	r.Get("/", h.listBikes)

	return r
}

func (h handler) listBikes(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	userID := ctx.Value(middleware.UserIDKey).(uuid.UUID)

	bikes, err := h.service.GetAllBikes(ctx, userID)
	if err != nil {
		return
	}

	render.Json(writer, mapToBikesJsonResponse(bikes))
}
