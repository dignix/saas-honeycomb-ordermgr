package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/domain"
)

type OrderRepository interface {
	Get(ctx context.Context, id uuid.UUID) (domain.Order, error)
	List(ctx context.Context) ([]domain.Order, error)
	Add(ctx context.Context, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Edit(ctx context.Context, id uuid.UUID) (domain.Order, error)
}
