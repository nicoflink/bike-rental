package rent

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

// Bike in terms of a rent.
// RentedBy (Optional) can be nil in case the bike is not rented.
type Bike struct {
	ID       uuid.UUID
	Location geo.Coordinates
	RentedBy *uuid.UUID
}

// startRent changes the status of the bike to "rented" by setting the value RentedBy to the provided renter.
func (b *Bike) startRent(renter uuid.UUID) error {
	if b.RentedBy != nil {
		return errors.New("bike is already rented")
	}

	b.RentedBy = &renter

	return nil
}

// stopRent changes the status of the bike to "available" by removing the renter.
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
