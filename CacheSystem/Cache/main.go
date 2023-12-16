package Cache

import (
	"fmt"
	"lld_problems/CacheSystem/Cache/EvictionPolicy"
	"lld_problems/CacheSystem/Cache/Storage"
)

type Cache struct {
	storage        Storage.Storage
	evictionPolicy EvictionPolicy.EvictionPolicy
}

func NewCache(storage Storage.Storage, policy EvictionPolicy.EvictionPolicy) *Cache {
	return &Cache{
		storage:        storage,
		evictionPolicy: policy,
	}
}

func (cache *Cache) Put(key string, value interface{}) error {
	err := cache.storage.Add(key, value)

	if err != nil {
		if err == Storage.FullStorageErr {
			evictedKey := cache.evictionPolicy.Evict()
			fmt.Println("Cache storage is full, trying to evict key: ", evictedKey)
			if err1 := cache.storage.Delete(evictedKey); err1 != nil {
				fmt.Println("error while evicting key: ", evictedKey)
				return err1
			}
		} else {
			return err
		}
	}

	if err = cache.storage.Add(key, value); err != nil {
		return err
	}
	cache.evictionPolicy.KeyAccessed(key)
	return nil
}

func (cache *Cache) Get(key string) (interface{}, error) {
	value, err := cache.storage.Get(key)
	cache.evictionPolicy.KeyAccessed(key)
	return value, err
}
