package rent

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

type Status uint8

func (s Status) Value() uint8 {
	return uint8(s)
}

const (
	Unknown Status = iota
	Started
	Finished
)

type Rent struct {
	ID            uuid.UUID
	Bike          uuid.UUID
	Renter        uuid.UUID
	Status        Status
	StartTime     time.Time
	EndTime       *time.Time
	StartLocation geo.Coordinates
	EndLocation   *geo.Coordinates
}

type StartRequest struct {
	Bike   uuid.UUID
	Renter uuid.UUID
}

type StopRequest struct {
	UserID uuid.UUID
	RentID uuid.UUID
}

type GetOpenRentsRequest struct {
	UserID uuid.UUID
}

func NewRent(bikedID uuid.UUID, renter uuid.UUID, location geo.Coordinates) *Rent {
	return &Rent{
		ID:            uuid.New(),
		Bike:          bikedID,
		Renter:        renter,
		Status:        Started,
		StartTime:     time.Now(),
		StartLocation: location,
	}
}

func (r *Rent) StopRent(renter uuid.UUID, endLocation geo.Coordinates) error {
	if r.Renter != renter {
		return errors.New("user not allowed to finish rent")
	}

	t := time.Now()
	r.EndTime = &t
	r.EndLocation = &endLocation
	r.Status = Finished

	return nil
}
