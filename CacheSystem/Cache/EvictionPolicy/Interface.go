package EvictionPolicy

type EvictionPolicy interface {
	KeyAccessed(key string)
	Evict() string
}
