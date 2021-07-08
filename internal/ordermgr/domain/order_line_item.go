package domain

import "github.com/google/uuid"

type OrderLineItem struct {
	ID       uuid.UUID
	Name     string
	Price    float32
	Quantity int
}
