package rent

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

type Bike struct {
	ID       uuid.UUID
	Location geo.Coordinates
	RentedBy *uuid.UUID
}

func (b *Bike) startRent(renter uuid.UUID) error {
	if b.RentedBy != nil {
		return errors.New("bike is already rented")
	}

	b.RentedBy = &renter

	return nil
}

func (b *Bike) stopRent(renter uuid.UUID) error {
	if b.RentedBy == nil {
		return nil
	}

	if *b.RentedBy != renter {
		return errors.New("user not allowed to return this bike")
	}

	b.RentedBy = nil

	return nil
}
