package cmd



type hotelReservation struct {
	reservationDate string
	room string
}

func (h hotelReservation) GetReservationDate() string {
	return h.reservationDate
}

func (h hotelReservation) CalculateCancellationFee() float64 {
	return 1.0
}


func (h *hotelReservation) ChangeType( name string) {
	h.room = name
}
