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

func mapRentToJsonResponse(r rent.Rent) RentResponse {
	var endTimeString string
	var endLocationString Coordinates

	if r.EndTime != nil {
		endTimeString = r.EndTime.String()
	}

	if r.EndLocation != nil {
		endLocationString = mapLocationToJson(*r.EndLocation)
	}

	return RentResponse{
		ID:            r.ID.String(),
		BikeID:        r.Bike.String(),
		Renter:        r.Renter.String(),
		StartTime:     r.StartTime.Format(time.RFC3339),
		EndTime:       endTimeString,
		StartLocation: mapLocationToJson(r.StartLocation),
		EndLocation:   endLocationString,
	}
}

func mapRentRequestToDomain(r RentRequest) rent.Request {
	return rent.Request{
		Bike:   uuid.MustParse(r.BikeID),
		Renter: uuid.MustParse(r.Renter),
	}
}
