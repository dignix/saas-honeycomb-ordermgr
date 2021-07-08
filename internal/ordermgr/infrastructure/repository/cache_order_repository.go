package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/domain"
)

type cacheOrderRepository struct {
	client *redis.Client
}

func NewCacheOrderRepository(client *redis.Client) *cacheOrderRepository {
	return &cacheOrderRepository{
		client: client,
	}
}

func (p *cacheOrderRepository) Get(ctx context.Context, ID uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (p *cacheOrderRepository) List(ctx context.Context) ([]domain.Order, error) {
	return []domain.Order{}, nil
}

func (p *cacheOrderRepository) Add(ctx context.Context, order domain.Order) (domain.Order, error) {
	return domain.Order{}, nil
}

func (p *cacheOrderRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	return nil
}

func (p *cacheOrderRepository) Edit(ctx context.Context, ID uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}
