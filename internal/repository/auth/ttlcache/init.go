package ttlcache

import (
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/iqbvl/lauk/internal/repository/auth"
)

func InitTTLCache() *ttlcache.Cache {
	return ttlcache.NewCache()
}

type Cache struct {
	Cache *ttlcache.Cache
}

func NewTTLCache(c *ttlcache.Cache) auth.Cache {
	return &Cache{
		Cache: c,
	}
}
