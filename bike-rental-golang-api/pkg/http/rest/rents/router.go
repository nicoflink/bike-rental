package rents

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/render"
)

type handler struct {
	service   ports.RentService
	validator ports.Validator
}

func NewRentsRouter(rentService ports.RentService, validator ports.Validator) chi.Router {
	r := chi.NewRouter()

	h := handler{
		service:   rentService,
		validator: validator,
	}

	r.Post("/", h.startRent)
	r.Patch("/{rentID}", h.stopRent)

	return r
}

func (h handler) startRent(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var startRequest StartRequest
	if err := jDecoder.Decode(&startRequest); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if err := h.validator.Struct(startRequest); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	domainRequest := mapStartRequestToDomain(startRequest)

	res, err := h.service.StartRent(ctx, domainRequest)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.Json(writer, mapRentToJsonResponse(res))
}

func (h handler) stopRent(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	userID := ctx.Value(middleware.UserIDKey).(uuid.UUID)

	rentStringID := chi.URLParam(request, "rentID")

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var stopRequest StopRequest
	if err := jDecoder.Decode(&stopRequest); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	stopRequest.UserID = userID
	stopRequest.RentID = rentStringID

	if err := h.validator.Struct(stopRequest); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	domainStopRequest := mapStopRequestToDomain(stopRequest)

	res, err := h.service.StopRent(ctx, domainStopRequest)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.Json(writer, mapRentToJsonResponse(res))
}
