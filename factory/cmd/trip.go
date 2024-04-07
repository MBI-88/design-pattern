package cmd



type trip struct {
	reservations []Reservation
}

func (t *trip) CalculateCancellationFee() float64 {
	total := 0.0

	for _, r := range t.reservations {
		total += r.CalculateCancellationFee()	
	}

	return total
}


func (t *trip) AddReservation(r Reservation) {
	t.reservations = append(t.reservations, r)
}