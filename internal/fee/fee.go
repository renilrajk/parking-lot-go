package fee

type Fee interface {
	RatePerHour() int
}

type fee struct {
	ratePerHour int
}

func NewFee(ratePerHour int) Fee {
	return &fee{
		ratePerHour: ratePerHour,
	}
}

func (f *fee) RatePerHour() int {
	return f.RatePerHour()
}
