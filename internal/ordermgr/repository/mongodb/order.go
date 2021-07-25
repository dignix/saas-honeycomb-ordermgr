package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr"
)

type mongoDBOrderRepository struct {
	dbHandle *mongo.Client
}

func NewMongoDBOrderRepository(dbHandle *mongo.Client) *mongoDBOrderRepository {
	return &mongoDBOrderRepository{
		dbHandle: dbHandle,
	}
}

func (r *mongoDBOrderRepository) FindByID(ctx context.Context, id string) (*ordermgr.Order, error) {
	return nil, nil
}

func (r *mongoDBOrderRepository) Save(ctx context.Context, order *ordermgr.Order) (*ordermgr.Order, error) {
	return nil, nil
}
