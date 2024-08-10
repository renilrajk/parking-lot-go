package spot

type Spot interface {
	Park()
	Unpark()
	IsAvailable() bool
}

type spot struct {
	isAvailable bool
}

func (s *spot) Park() {
	s.isAvailable = false
}

func (s *spot) Unpark() {
	s.isAvailable = true
}

func (s *spot) IsAvailable() bool {
	return s.isAvailable
}
