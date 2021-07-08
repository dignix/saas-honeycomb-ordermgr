package domain

import "github.com/google/uuid"

const (
	APPROVAL_PENDING = iota
	APPROVED
	REJECTED
	CANCEL_PENDING
	CANCELLED
	REVISION_PENDING
)

type Order struct {
	ID             uuid.UUID
	ConsumerID     uuid.UUID
	RestaurantID   uuid.UUID
	OrderLineItems OrderLineItem
	DeliveryInfo   DeliveryInfo
	PaymentInfo    PaymentInfo
	State          int
}
