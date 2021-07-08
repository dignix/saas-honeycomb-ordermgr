package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/domain"
	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/domain/repository"
)

type OrderService interface {
	List(ctx context.Context) ([]domain.Order, error)
	Create(
		ctx context.Context,
		consumerID uuid.UUID,
		restaurantID uuid.UUID,
		deliveryInfo domain.DeliveryInfo,
		lineItems []domain.OrderLineItem,
	) (domain.Order, error)
	Cancel(ctx context.Context, id uuid.UUID) (domain.Order, error)
	Approve(ctx context.Context, id uuid.UUID) (domain.Order, error)
	Reject(ctx context.Context, id uuid.UUID) (domain.Order, error)
}

type DomainOrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *DomainOrderService {
	return &DomainOrderService{
		orderRepository: orderRepository,
	}
}

func (d *DomainOrderService) List(ctx context.Context) ([]domain.Order, error) {
	orders, err := d.orderRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (d *DomainOrderService) Create(
	ctx context.Context,
	consumerID uuid.UUID,
	restaurantID uuid.UUID,
	deliveryInfo domain.DeliveryInfo,
	lineItems []domain.OrderLineItem,
) (domain.Order, error) {
	return domain.Order{}, nil
}

func (d *DomainOrderService) Cancel(ctx context.Context, id uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (d *DomainOrderService) Approve(ctx context.Context, id uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (d *DomainOrderService) Reject(ctx context.Context, id uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}
