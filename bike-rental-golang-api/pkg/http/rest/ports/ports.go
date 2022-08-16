package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/list"
	"github.com/nicoflink/bike-rental/pkg/rent"
)

type ListService interface {
	GetAllBikes(ctx context.Context, userID uuid.UUID) ([]list.Bike, error)
	UpdateBikePosition(ctx context.Context, userID uuid.UUID, locationUpdate list.BikeLocationUpdate) (list.Bike, error)
}

type RentService interface {
	StartRent(ctx context.Context, start rent.StartRequest) (rent.Rent, error)
	StopRent(ctx context.Context, stop rent.StopRequest) (rent.Rent, error)
}

type Validator interface {
	Struct(s interface{}) error
}
