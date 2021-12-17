package repository

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
)

type Cache interface {
	GetUser(ctx context.Context, key string) (model.User, error)
	SetUser(ctx context.Context, in model.User) error
	GetRates(ctx context.Context, key string) (float64, error)
	SetRates(ctx context.Context, key string, rates float64) error
}
