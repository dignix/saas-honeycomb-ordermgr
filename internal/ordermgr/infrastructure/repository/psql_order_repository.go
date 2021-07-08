package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/domain"
)

type psqlOrderRepository struct {
	dbHandle *sql.DB
}

func NewPSQLOrderRepository(dbHandle *sql.DB) *psqlOrderRepository {
	return &psqlOrderRepository{
		dbHandle: dbHandle,
	}
}

func (p *psqlOrderRepository) Get(ctx context.Context, ID uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (p *psqlOrderRepository) List(ctx context.Context) ([]domain.Order, error) {
	return []domain.Order{}, nil
}

func (p *psqlOrderRepository) Add(ctx context.Context, order domain.Order) (domain.Order, error) {
	return domain.Order{}, nil
}

func (p *psqlOrderRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	return nil
}

func (p *psqlOrderRepository) Edit(ctx context.Context, ID uuid.UUID) (domain.Order, error) {
	return domain.Order{}, nil
}
