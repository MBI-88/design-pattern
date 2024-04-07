package cmd 



type flightReservation struct {
	reservationDate string
	luggageAllowed  int
}

func (f flightReservation) AddExtraLuggageAllowance(peices int) {
	f.luggageAllowed = peices
}

func (f flightReservation) CalculateCancellationFee() float64 {
	return 2.0
}

func (f flightReservation) GetReservationDate() string {
	return f.reservationDate
}
