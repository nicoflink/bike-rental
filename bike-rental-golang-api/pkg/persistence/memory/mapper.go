package memory

import (
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/list"
	"github.com/nicoflink/bike-rental/pkg/rent"
)

// mapBikeToListBike maps persistence bike to list bike.
// In case RentedByUser is set - the bike is rented.
// In case RentedByUser is set and matches the provided user, rentedByUser is true.
func mapBikeToListBike(b Bike, userID uuid.UUID) list.Bike {
	var rentedByUser bool
	var rented bool

	if b.RentedByUser != nil {
		rented = true

		if *b.RentedByUser == userID {
			rentedByUser = true
		}
	}

	return list.Bike{
		ID:   b.ID,
		Name: b.Name,
		Location: geo.Coordinates{
			Lat: b.Location.Lat,
			Lng: b.Location.Lng,
		},
		Rented:       rented,
		RentedByUser: rentedByUser,
	}
}

// mapBikeToRentBike maps persistence bike to rent bike.
func mapBikeToRentBike(b Bike) rent.Bike {
	return rent.Bike{
		ID:       b.ID,
		Location: b.Location,
		RentedBy: b.RentedByUser,
	}
}
