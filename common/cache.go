package common

import "github.com/patrickmn/go-cache"

var Cache *cache.Cache

func init() {
	Cache = NewCache()
}

func NewCache() *cache.Cache {
	return cache.New(cache.NoExpiration, cache.NoExpiration)
}
