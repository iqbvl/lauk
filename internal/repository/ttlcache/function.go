package ttlcache

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
	"github.com/iqbvl/lauk/internal/platform/util"
	log "github.com/sirupsen/logrus"
)

func (c *Cache) GetUser(ctx context.Context, key string) (model.User, error) {
	var u model.User
	data, err := c.Cache.Get(key)
	if err != nil {
		if err.Error() != model.KeyNotFound {
			log.Infof("[Cache][GetUser] msg %s \n", err.Error())
		} else {
			log.Errorf("[Cache][GetUser] msg %s \n", err.Error())
		}

		return u, err
	}

	u = data.(model.User)
	return u, nil
}

func (c *Cache) SetUser(ctx context.Context, in model.User) error {
	key := util.GenerateKey(in)
	err := c.Cache.Set(key, in)
	if err != nil {
		log.Errorf("[Cache][SetUser] msg %s \n", err.Error())
		return err
	}

	return nil
}

func (c *Cache) SetRates(ctx context.Context, key string, rates float64) error {
	err := c.Cache.SetWithTTL(key, rates, util.GetRatesExpiryTime())
	if err != nil {
		log.Errorf("[Cache][SetRates] msg %s \n", err.Error())
		return err
	}

	return nil
}

func (c *Cache) GetRates(ctx context.Context, key string) (float64, error) {
	data, err := c.Cache.Get(key)
	if err != nil {
		if err.Error() != model.KeyNotFound {
			log.Errorf("[Cache][GetRates] msg %s \n", err.Error())
			return float64(0), err
		}

		log.Infof("[Cache][GetRates] msg %s \n", err.Error())
		return float64(0), err
	}

	return data.(float64), nil
}
