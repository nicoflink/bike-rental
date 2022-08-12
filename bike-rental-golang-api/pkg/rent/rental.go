package rent

import (
	"time"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

type Rent struct {
	ID            uuid.UUID
	Bike          uuid.UUID
	Renter        uuid.UUID
	StartTime     time.Time
	EndTime       *time.Time
	StartLocation geo.Coordinates
	EndLocation   *geo.Coordinates
}

type Request struct {
	Bike   uuid.UUID
	Renter uuid.UUID
}

func NewRent(bikedID uuid.UUID, renter uuid.UUID, location geo.Coordinates) *Rent {
	return &Rent{
		ID:            uuid.New(),
		Bike:          bikedID,
		Renter:        renter,
		StartTime:     time.Now(),
		StartLocation: location,
	}
}

func (r *Rent) StopRent(endLocation geo.Coordinates) {
	t := time.Now()
	r.EndTime = &t
	r.EndLocation = &endLocation
}
