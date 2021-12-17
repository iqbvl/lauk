package ttlcache

import (
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/iqbvl/lauk/internal/repository"
) 

type Cache struct {
	Cache *ttlcache.Cache
}

func NewTTLCache(c *ttlcache.Cache) repository.Cache {
	return &Cache{
		Cache: c,
	}
}
