package rents

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/rent"
)

func mapLocationToJson(g geo.Coordinates) Coordinates {
	return Coordinates{
		Lat: fmt.Sprintf("%f", g.Lat),
		Lng: fmt.Sprintf("%f", g.Lng),
	}
}

func mapRentsToJsonResponse(r []rent.Rent) []RentResponse {
	rents := make([]RentResponse, 0, len(r))

	for _, ren := range r {
		rents = append(rents, mapRentToJsonResponse(ren))
	}

	return rents
}

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

func mapStartRequestToDomain(r StartRequest) rent.StartRequest {
	return rent.StartRequest{
		Bike:   uuid.MustParse(r.BikeID),
		Renter: uuid.MustParse(r.Renter),
	}
}

func mapStopRequestToDomain(request StopRequest) rent.StopRequest {
	return rent.StopRequest{
		UserID: request.UserID,
		RentID: uuid.MustParse(request.RentID),
	}
}

func mapGetRentRequestToDomain(request GetRentRequest) rent.GetOpenRentsRequest {
	return rent.GetOpenRentsRequest{UserID: uuid.MustParse(request.UserID)}
}
