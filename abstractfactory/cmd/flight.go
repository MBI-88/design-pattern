package cmd



type flightReservation struct {
	reservationDate string
	payment  float64
}

func (f *flightReservation) SetPayment(p float64) {
	f.payment = p
}

func (f flightReservation) CalculateCancellationFee() float64 {
	return 2.0
}

func (f flightReservation) GetReservationDate() string {
	return f.reservationDate
}

func (f *flightReservation) SetReservationDate( d string) {
	f.reservationDate = d
}