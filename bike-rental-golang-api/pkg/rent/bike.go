package rent

import (
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

type Bike struct {
	ID       uuid.UUID
	Location geo.Coordinates
	RentedBy *uuid.UUID
}

func (b *Bike) removeRenter() {
	b.RentedBy = nil
}
