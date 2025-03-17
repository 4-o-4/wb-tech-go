package cache

import (
    "L0/internal/domain"
    "sync"
)

type MemoryCache struct {
    data map[string]*domain.Order
    mux  sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
    return &MemoryCache{
        data: make(map[string]*domain.Order),
    }
}

func (c *MemoryCache) Set(key string, value *domain.Order) {
    c.mux.Lock()
    defer c.mux.Unlock()
    c.data[key] = value
}

func (c *MemoryCache) Get(key string) (*domain.Order, bool) {
    c.mux.RLock()
    defer c.mux.RUnlock()
    value, ok := c.data[key]
    return value, ok
}
