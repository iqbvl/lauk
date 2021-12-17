package fetch

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
)

func (u *FetchUsecase) GetStorages(ctx context.Context, r model.GetStoragesRequest) ([]model.Storage, error) {
	var res []model.Storage

	res, err := u.External.GetStorages(ctx, r)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *FetchUsecase) GetRates(ctx context.Context) (float64, error) {

	rates, err := u.TTLCache.GetRates(ctx, model.IDR_USD)
	if err != nil {
		if err.Error() != model.KeyNotFound {
			return float64(0), err
		} 
	}

	if rates == 0 {
		rates, err = u.External.GetRates(ctx)
		if err != nil {
			return float64(0), err
		}

		err = u.TTLCache.SetRates(ctx, model.IDR_USD, rates)
		if err != nil {
			return float64(0), err
		}
	}

	return rates, nil
}
