package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/iqbvl/lauk/internal/model"
	log "github.com/sirupsen/logrus"
)

func (c *External) GetStorages(ctx context.Context, r model.GetStoragesRequest) ([]model.Storage, error) {
	var s []model.Storage

	request, err := http.NewRequest(http.MethodGet, storageURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	// assign price USD
	var res []model.Storage
	for _, v := range s {
		var (
			fPrice  float64
			storage model.Storage
		)

		if v.Price == "" {
			continue
		}

		if fPrice, err = strconv.ParseFloat(v.Price, 64); err != nil {
			log.Infof("[repository][GetStorages] failed parse float : %s \n", err.Error())
			continue
		}

		storage = v
		storage.PriceUSD = fmt.Sprintf("%f", r.Rates*fPrice)
		res = append(res, storage)
	}

	return res, nil
}

func (c *External) GetRates(ctx context.Context) (float64, error) {
	var (
		result float64
		resp   Rates
	)

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(converterURL, model.IDR_USD, converterAPIKey), nil)
	if err != nil {
		return result, err
	}

	response, err := c.Client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return result, err
	}

	return resp.IDRUSD, nil
}
