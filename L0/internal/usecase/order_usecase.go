package usecase

import (
    "L0/internal/cache"
    "L0/internal/domain"
    "L0/internal/repository/postgres"
)

type OrderUseCase struct {
    repo  *postgres.OrderRepository
    cache *cache.MemoryCache
}

func NewOrderUseCase(r *postgres.OrderRepository, c *cache.MemoryCache) *OrderUseCase {
    return &OrderUseCase{repo: r, cache: c}
}

func (uc *OrderUseCase) GetOrder(id string) (*domain.Order, error) {
    if order, found := uc.cache.Get(id); found {
        return order, nil
    }
    order, err := uc.repo.GetOrderById(id)
    if err != nil {
        return nil, err
    }
    uc.cache.Set(id, order)
    return order, nil
}
