package ports

import (
	"context"

	"github.com/nicoflink/bike-rental/pkg/list"
)

type ListService interface {
	GetAllBikes(ctx context.Context) ([]list.Bike, error)
}

type Validator interface {
	Struct(s interface{}) error
}
