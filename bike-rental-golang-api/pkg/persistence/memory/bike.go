package memory

import (
	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

// Bike representation in the data layer.
type Bike struct {
	ID           uuid.UUID
	Name         string
	Location     geo.Coordinates
	RentedByUser *uuid.UUID
}
