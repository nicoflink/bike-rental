package rents

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	e "github.com/nicoflink/bike-rental/pkg/http/rest/errors"
	"github.com/nicoflink/bike-rental/pkg/http/rest/middleware"
	"github.com/nicoflink/bike-rental/pkg/http/rest/ports"
	"github.com/nicoflink/bike-rental/pkg/http/rest/render"
)

// handler is an internal struct providing different handler functions for the specific routes.
// It is responsible to validate the incoming request and call the corresponding domain service.
type handler struct {
	service   ports.RentService
	validator ports.Validator
}

// NewRentsRouter returns a router for rents resource.
func NewRentsRouter(rentService ports.RentService, validator ports.Validator) (chi.Router, error) {
	if rentService == nil || validator == nil {
		return nil, errors.New("NewRentsRouter dependencies are not fulfilled")
	}

	r := chi.NewRouter()

	h := handler{
		service:   rentService,
		validator: validator,
	}

	r.Get("/", h.getRents)
	r.Post("/", h.startRent)
	r.Patch("/{rentID}", h.stopRent)

	return r, nil
}

// startRent handles the request to start a rent.
func (h handler) startRent(writer http.ResponseWriter, request *http.Request) {
	const prefix = "rents.handler.startRent"

	ctx := request.Context()

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, e.ErrEmptyRequestBody, http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var startRequest StartRequest
	if err := jDecoder.Decode(&startRequest); err != nil {
		http.Error(writer, e.ErrParsingJson, http.StatusInternalServerError)

		return
	}

	if err := h.validator.Struct(startRequest); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	domainRequest := mapStartRequestToDomain(startRequest)

	res, err := h.service.StartRent(ctx, domainRequest)
	if err != nil {
		log.Println(fmt.Sprintf("%s: ERROR - %v", prefix, err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.ToJson(writer, mapRentToJsonResponse(res))
}

// stopRent handles the request to stop a rent.
func (h handler) stopRent(writer http.ResponseWriter, request *http.Request) {
	const prefix = "rents.handler.stopRent"

	ctx := request.Context()
	userID := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	rentStringID := chi.URLParam(request, "rentID")

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, e.ErrEmptyRequestBody, http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var stopRequest StopRequest
	if err := jDecoder.Decode(&stopRequest); err != nil {
		http.Error(writer, e.ErrParsingJson, http.StatusInternalServerError)

		return
	}

	stopRequest.UserID = userID
	stopRequest.RentID = rentStringID

	if err := h.validator.Struct(stopRequest); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	domainStopRequest := mapStopRequestToDomain(stopRequest)

	res, err := h.service.StopRent(ctx, domainStopRequest)
	if err != nil {
		log.Println(fmt.Sprintf("%s: ERROR - %v", prefix, err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.ToJson(writer, mapRentToJsonResponse(res))
}

// getRents handles the request to get all started rents of a user.
func (h handler) getRents(writer http.ResponseWriter, request *http.Request) {
	const prefix = "rents.handler.getRents"

	ctx := request.Context()

	queryValues := request.URL.Query()
	if len(queryValues) == 0 {
		http.Error(writer, e.ErrGetRentQuery, http.StatusBadRequest)

		return
	}

	status := queryValues.Get("status")
	userID := queryValues.Get("userID")

	req := GetRentRequest{
		Status: status,
		UserID: userID,
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	domainGetRequest := mapGetRentRequestToDomain(req)

	res, err := h.service.GetStartedRents(ctx, domainGetRequest)
	if err != nil {
		log.Println(fmt.Sprintf("%s: ERROR - %v", prefix, err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.ToJson(writer, mapRentsToJsonResponse(res))
}
