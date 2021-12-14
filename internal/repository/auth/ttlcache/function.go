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
		if err.Error() != "key not found" {
			log.Infof("[Cache][GetUser] msg %s", err.Error())
		} else {
			log.Errorf("[Cache][GetUser] msg %s", err.Error())
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
		log.Errorf("[Cache][SetUser] msg %s", err.Error())
		return err
	}

	return nil
}
