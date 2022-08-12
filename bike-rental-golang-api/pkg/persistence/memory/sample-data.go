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
		// Bikes available
		{
			ID:   uuid.MustParse("c07e1166-4971-437a-879b-44a11c8f45b2"),
			Name: "Carlo",
			Location: geo.Coordinates{
				Lat: 50.10753883916688,
				Lng: 8.656987445600874,
			},
			RentedByUser: nil,
		},
		{
			ID:   uuid.MustParse("4aec8d9c-938e-49bb-b301-944fa991319c"),
			Name: "Linus",
			Location: geo.Coordinates{
				Lat: 50.10150556308182,
				Lng: 8.660624044815425,
			},
			RentedByUser: nil,
		},
		{
			ID:   uuid.MustParse("2f7521b4-983f-47de-98ba-f29c6da2f93e"),
			Name: "Jette",
			Location: geo.Coordinates{
				Lat: 50.106680355676026,
				Lng: 8.676556755764615,
			},
			RentedByUser: nil,
		},
		{
			ID:   uuid.MustParse("67123f76-8751-463c-b72a-23abe633a246"),
			Name: "Alva",
			Location: geo.Coordinates{
				Lat: 50.11070702389974,
				Lng: 8.660188011071721,
			},
			RentedByUser: uuidToPointer(uuid.MustParse("10debbdb-1a92-4128-bdb0-cc381ea5585f")),
		},
		// Bike already rented by unknown user
		{
			ID:   uuid.MustParse("a85ea086-72f3-4f4d-b25d-83a43e995206"),
			Name: "Nina",
			Location: geo.Coordinates{
				Lat: 50.1123635233249,
				Lng: 8.649741801640113,
			},
			RentedByUser: uuidToPointer(uuid.New()),
		},
	}
)
