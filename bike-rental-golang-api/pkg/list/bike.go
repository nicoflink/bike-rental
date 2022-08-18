package list

import (
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

// Bike is bike used for the view in the bike-app
// Rented: Bike is rented
// RentedByUser: Bike is rented by the current user requesting the resource.
type Bike struct {
	ID           uuid.UUID
	Name         string
	Location     geo.Coordinates
	Rented       bool
	RentedByUser bool
}

// BikeLocationUpdate used for location update of the bike.
type BikeLocationUpdate struct {
	BikeID   uuid.UUID
	Location geo.Coordinates
}
