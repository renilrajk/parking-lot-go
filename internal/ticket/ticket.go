package ticket

import (
	"time"

	"github.com/renilrajk/parking-lot-go/internal/spot"
)

type Ticket interface {
	TotalDuration() time.Duration
	GetSpot() spot.Spot
}

type ticket struct {
	entryTime     time.Time
	vehicleNumber string
	spot          spot.Spot
}

func NewTicket(vehicleNumber string, sp spot.Spot) Ticket {
	return &ticket{
		entryTime:     time.Now(),
		vehicleNumber: vehicleNumber,
		spot:          sp,
	}
}

func (t *ticket) TotalDuration() time.Duration {
	now := time.Now()
	return now.Sub(t.entryTime)
}

func (t *ticket) GetSpot() spot.Spot {
	return t.spot
}
