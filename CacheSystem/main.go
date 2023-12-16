package main

import (
	"fmt"
	"lld_problems/CacheSystem/Cache"
	"lld_problems/CacheSystem/Cache/EvictionPolicy"
	"lld_problems/CacheSystem/Cache/Storage"
	"strconv"
)

func main() {
	storage := Storage.NewMapStorage(5)
	evictionPolicy := EvictionPolicy.NewLRUEvictionPolicy()
	cache := Cache.NewCache(storage, evictionPolicy)
	for i := 1; i <= 5; i++ {
		cache.Put(strconv.Itoa(i), i)
	}
	cache.Get(strconv.Itoa(2))
	cache.Get(strconv.Itoa(5))
	cache.Get(strconv.Itoa(1))
	cache.Get(strconv.Itoa(4))
	cache.Get(strconv.Itoa(1))
	cache.Put(strconv.Itoa(6), 6)
	for i := 1; i <= 6; i++ {
		value, _ := cache.Get(strconv.Itoa(i))
		fmt.Println(value)
	}

}
