package cmd

type PremiumReservation struct {
	Reservation
}

func (p PremiumReservation) Cancel() {
	p.sellerRef.CancelReservation(0)
}
