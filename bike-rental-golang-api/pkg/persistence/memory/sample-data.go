package memory

import (
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

func uuidToPointer(u uuid.UUID) *uuid.UUID {
	return &u
}

var (
	SampleBikes = []Bike{
		{
			ID:   uuid.New(),
			Name: "Carlo",
			Location: geo.Coordinates{
				Lat: 50.10753883916688,
				Lng: 8.656987445600874,
			},
			RentedByUser: nil,
		},
		{
			ID:   uuid.New(),
			Name: "Alva",
			Location: geo.Coordinates{
				Lat: 50.11070702389974,
				Lng: 8.660188011071721,
			},
			RentedByUser: uuidToPointer(uuid.New()),
		},
		{
			ID:   uuid.New(),
			Name: "Nina",
			Location: geo.Coordinates{
				Lat: 50.1123635233249,
				Lng: 8.649741801640113,
			},
			RentedByUser: nil,
		},
	}
)
