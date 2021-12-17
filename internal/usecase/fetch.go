package usecase

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
)

type Fetch interface {
	GetStorages(ctx context.Context, r model.GetStoragesRequest) ([]model.Storage, error)
	GetRates(ctx context.Context) (float64, error)
}
