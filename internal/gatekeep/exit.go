package gatekeep

import "github.com/renilrajk/parking-lot-go/internal/spot"

type Exit interface {
	Exit()
}

type exit struct {
}

func (e *exit) Exit(spot spot.Spot) {
	spot.Unpark()
}
