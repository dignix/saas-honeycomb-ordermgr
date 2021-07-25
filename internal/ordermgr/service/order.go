package service

import (
	"context"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr"
)

type OrderRepository interface {
	FindByID(ctx context.Context, id string) (*ordermgr.Order, error)
	Save(ctx context.Context, order *ordermgr.Order) (*ordermgr.Order, error)
}

type orderService struct {
	orderRepository OrderRepository
}

func NewOrderService(orderRepository OrderRepository) *orderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (s *orderService) Get(ctx context.Context, id string) (*ordermgr.Order, error) {
	order, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) Create(ctx context.Context, order *ordermgr.Order) (*ordermgr.Order, error) {
	order, err := s.orderRepository.Save(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
