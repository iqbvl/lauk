package repository

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
)

type External interface {
	GetStorages(ctx context.Context, r model.GetStoragesRequest) ([]model.Storage, error)
	GetRates(ctx context.Context) (float64, error)
}
