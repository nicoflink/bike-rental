package bikes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/render"
	"github.com/nicoflink/bike-rental/pkg/list"
)

type handler struct {
	service   ports.ListService
	validator ports.Validator
}

func NewBikesRouter(listService ports.ListService, validator ports.Validator) chi.Router {
	r := chi.NewRouter()

	h := handler{
		service:   listService,
		validator: validator,
	}

	r.Get("/", h.listBikes)
	r.Patch("/{bikeID}", h.updateBikeLocation)

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

func (h handler) updateBikeLocation(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	userID := ctx.Value(middleware.UserIDKey).(uuid.UUID)

	bikeStringID := chi.URLParam(request, "bikeID")

	bikeID, err := uuid.Parse(bikeStringID)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var locationUpdate BikeLocationUpdate
	if err := jDecoder.Decode(&locationUpdate); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if err := h.validator.Struct(locationUpdate); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	res, err := h.service.UpdateBikePosition(ctx, userID, list.BikeLocationUpdate{
		BikeID:   bikeID,
		Location: geo.Coordinates(locationUpdate.Location),
	})
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.Json(writer, mapToBikeJsonResponse(res))
}
