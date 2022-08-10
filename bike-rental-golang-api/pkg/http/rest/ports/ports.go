package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/list"
)

type ListService interface {
	GetAllBikes(ctx context.Context, userID uuid.UUID) ([]list.Bike, error)
}

type Validator interface {
	Struct(s interface{}) error
}
