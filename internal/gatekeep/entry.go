package gatekeep

import (
	"github.com/renilrajk/parking-lot-go/internal/spot"
)

type Entry interface {
	Enter()
}

type entry struct {
}

func (e *entry) Enter(spot spot.Spot) {
	spot.Park()
}
