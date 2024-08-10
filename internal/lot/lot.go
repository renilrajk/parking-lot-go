package lot

import (
	"errors"

	"github.com/renilrajk/parking-lot-go/internal/fee"
	"github.com/renilrajk/parking-lot-go/internal/spot"
	"github.com/renilrajk/parking-lot-go/internal/ticket"
)

type Lot interface {
	Enter(vehicleNumber string) (ticket.Ticket, error)
	Exit(ticket ticket.Ticket) float64
}

type lot struct {
	total    int
	occupied int
	spots    []spot.Spot
	fee      fee.Fee
}

func (l *lot) Enter(vehicleNumber string) (ticket.Ticket, error) {
	if l.occupied < l.total {
		s := l.getNearestSpot()
		ticket := ticket.NewTicket(vehicleNumber, s)

		s.Park()
		l.occupied++

		return ticket, nil
	}
	return nil, errors.New("Parking lot is full")
}

func (l *lot) Exit(ticket ticket.Ticket) float64 {
	fee := ticket.TotalDuration().Hours() * float64(l.fee.RatePerHour())
	ticket.GetSpot().Unpark()
	l.occupied--

	return fee
}

func (l *lot) getNearestSpot() spot.Spot {
	for _, s := range l.spots {
		if s.IsAvailable() {
			return s
		}
	}
	return nil
}
