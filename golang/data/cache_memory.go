package data

import (
	"fmt"
	"time"
)

type (
	CacheMemory struct {
		values map[string]*CacheItem
	}
	CacheItem struct {
		value      interface{}
		validUntil time.Time
	}
)

func (c *CacheMemory) Setup(args ...interface{}) *CacheMemory {
	c.values = make(map[string]*CacheItem)
	return c
}

func (c *CacheMemory) Set(key string, value interface{}, timeToLive time.Duration) error {
	c.values[key] = &CacheItem{
		value:      value,
		validUntil: time.Now().Add(timeToLive),
	}
	return nil
}

func (c *CacheMemory) Get(key string) (interface{}, error) {
	if value, ok := c.values[key]; ok {
		if value.validUntil.After(time.Now()) {
			return value.value, nil
		}
		return nil, fmt.Errorf("Expired value for key %s", key)
	}
	return nil, fmt.Errorf("Key not found %s", key)
}
