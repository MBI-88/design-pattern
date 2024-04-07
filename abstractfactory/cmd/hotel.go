package cmd



type hotelReservation struct {
	reservationDate string
	payment float64
}

func (h hotelReservation) GetReservationDate() string {
	return h.reservationDate
}

func (h hotelReservation) CalculateCancellationFee() float64 {
	return 1.0
}

func (h *hotelReservation) SetReservationDate( d string) {
	h.reservationDate = d
}

func (h *hotelReservation) SetPayment(p float64) {
	h.payment = p
}

