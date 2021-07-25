package ordermgr

import "github.com/google/uuid"

type DeliveryInfo struct {
	ID      uuid.UUID
	Time    string
	Address string
}
