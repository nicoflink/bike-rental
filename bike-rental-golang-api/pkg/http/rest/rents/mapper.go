package rents

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/rent"
)

// mapLocationToJson maps coordinates to strings
func mapLocationToJson(g geo.Coordinates) Coordinates {
	return Coordinates{
		Lat: fmt.Sprintf("%f", g.Lat),
		Lng: fmt.Sprintf("%f", g.Lng),
	}
}

// mapRentsToJsonResponse maps domain rents to json response.
// If the rent slice is empty, the response also returns an empty slice.
func mapRentsToJsonResponse(r []rent.Rent) []RentResponse {
	rents := make([]RentResponse, 0, len(r))

	for _, ren := range r {
		rents = append(rents, mapRentToJsonResponse(ren))
	}

	return rents
}

// mapRentToJsonResponse maps a single domain rent to the JSON response.
func mapRentToJsonResponse(r rent.Rent) RentResponse {
	const isoFormat = time.RFC3339

	var endTimeString string
	var endLocationString Coordinates

	if r.EndTime != nil {
		endTimeString = r.EndTime.Format(isoFormat)
	}

	if r.EndLocation != nil {
		endLocationString = mapLocationToJson(*r.EndLocation)
	}

	return RentResponse{
		ID:            r.ID.String(),
		BikeID:        r.Bike.String(),
		Renter:        r.Renter.String(),
		Status:        r.Status.Value(),
		StartTime:     r.StartTime.Format(isoFormat),
		EndTime:       endTimeString,
		StartLocation: mapLocationToJson(r.StartLocation),
		EndLocation:   endLocationString,
	}
}

// mapStartRequestToDomain maps json StartRequest to rent StartRequest.
// As json request is already validated by the handler, MustParse of the uuid pkg can be used.
func mapStartRequestToDomain(r StartRequest) rent.StartRequest {
	return rent.StartRequest{
		Bike:   uuid.MustParse(r.BikeID),
		Renter: uuid.MustParse(r.Renter),
	}
}

// mapStopRequestToDomain maps json StopRequest to rent StopRequest.
// As json request is already validated by the handler, MustParse of the uuid pkg can be used.
// The UserID of the json StopRequest is extracted by the middleware and therefor already validated.
func mapStopRequestToDomain(request StopRequest) rent.StopRequest {
	return rent.StopRequest{
		UserID: request.UserID,
		RentID: uuid.MustParse(request.RentID),
	}
}

// mapGetRentRequestToDomain maps json GetRentRequest to rent GetOpenRentsRequest.
// As json request is already validated by the handler, MustParse of the uuid pkg can be used.
func mapGetRentRequestToDomain(request GetRentRequest) rent.GetOpenRentsRequest {
	return rent.GetOpenRentsRequest{UserID: uuid.MustParse(request.UserID)}
}
