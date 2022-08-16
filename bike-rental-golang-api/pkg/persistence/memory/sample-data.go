package memory

import (
	"time"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/rent"
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
			RentedByUser: uuidToPointer(uuid.MustParse("d2946438-030d-4350-ae5c-e5b4f8dc402d")),
		},
	}
	SampleRents = []rent.Rent{
		{
			ID:        uuid.MustParse("6733d0dc-4138-4bb8-a68f-79751674fc96"),
			Bike:      uuid.MustParse("67123f76-8751-463c-b72a-23abe633a246"),
			Renter:    uuid.MustParse("10debbdb-1a92-4128-bdb0-cc381ea5585f"),
			Status:    rent.Started,
			StartTime: time.Now(),
			EndTime:   nil,
			StartLocation: geo.Coordinates{
				Lat: 50.11070702389974,
				Lng: 8.660188011071721,
			},
			EndLocation: nil,
		},
		{
			ID:        uuid.MustParse("8772940e-4f2a-454d-8ab2-b3a3b935ca7c"),
			Bike:      uuid.MustParse("a85ea086-72f3-4f4d-b25d-83a43e995206"),
			Renter:    uuid.MustParse("d2946438-030d-4350-ae5c-e5b4f8dc402d"),
			Status:    rent.Started,
			StartTime: time.Now(),
			EndTime:   nil,
			StartLocation: geo.Coordinates{
				Lat: 50.1123635233249,
				Lng: 8.649741801640113,
			},
			EndLocation: nil,
		},
	}
)
