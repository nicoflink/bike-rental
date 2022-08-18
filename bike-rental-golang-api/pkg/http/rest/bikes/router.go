package bikes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	e "github.com/nicoflink/bike-rental/pkg/http/rest/errors"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/render"
	"github.com/nicoflink/bike-rental/pkg/list"
)

// handler is an internal struct providing different handler functions for the specific routes.
// It is responsible to validate the incoming request and call the corresponding domain service.
type handler struct {
	service   ports.ListService
	validator ports.Validator
}

// NewBikesRouter returns a router for bikes resource.
func NewBikesRouter(listService ports.ListService, validator ports.Validator) (chi.Router, error) {
	if listService == nil || validator == nil {
		return nil, errors.New("NewBikesRouter dependencies are not fulfilled")
	}

	r := chi.NewRouter()

	h := handler{
		service:   listService,
		validator: validator,
	}

	r.Get("/", h.listBikes)
	r.Patch("/{bikeID}", h.updateBikeLocation)

	return r, nil
}

// listBikes list all bikes in the context of the current user.
func (h handler) listBikes(writer http.ResponseWriter, request *http.Request) {
	const prefix = "bikes.handler.listBikes"

	ctx := request.Context()

	userID := ctx.Value(middleware.UserIDKey).(uuid.UUID)

	bikes, err := h.service.GetAllBikes(ctx, userID)
	if err != nil {
		log.Println(fmt.Sprintf("%s: ERROR - %v", prefix, err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.ToJson(writer, mapToBikesJsonResponse(bikes))
}

// updateBikeLocation handles the request to update the position of a bike.
func (h handler) updateBikeLocation(writer http.ResponseWriter, request *http.Request) {
	const prefix = "bikes.handler.updateBikeLocation"

	ctx := request.Context()

	userID := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	bikeStringID := chi.URLParam(request, "bikeID")

	err := h.validator.Var(bikeStringID, "required,uuid4")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, e.ErrEmptyRequestBody, http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var locationUpdate BikeLocationUpdate
	if lErr := jDecoder.Decode(&locationUpdate); lErr != nil {
		http.Error(writer, e.ErrParsingJson, http.StatusInternalServerError)

		return
	}

	if lErr := h.validator.Struct(locationUpdate); lErr != nil {
		http.Error(writer, lErr.Error(), http.StatusBadRequest)

		return
	}

	res, err := h.service.UpdateBikePosition(ctx, userID, list.BikeLocationUpdate{
		BikeID:   uuid.MustParse(bikeStringID),
		Location: geo.Coordinates(locationUpdate.Location),
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s: ERROR - %v", prefix, err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.ToJson(writer, mapToBikeJsonResponse(res))
}
