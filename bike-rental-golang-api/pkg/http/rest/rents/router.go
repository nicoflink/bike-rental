package rents

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	r.Post("/", h.createRent)

	return r
}

func (h handler) createRent(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	rBody := request.Body
	if rBody == nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	jDecoder := json.NewDecoder(rBody)
	jDecoder.DisallowUnknownFields()

	var rentRequest RentRequest
	if err := jDecoder.Decode(&rentRequest); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if err := h.validator.Struct(rentRequest); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	domainRequest := mapRentRequestToDomain(rentRequest)

	res, err := h.service.StartRent(ctx, domainRequest)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	render.Json(writer, mapRentToJsonResponse(res))
}
