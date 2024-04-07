package cmd

type Reservation struct {
	sellerRef SellerInt
}

func (r Reservation) Cancel() {
	r.sellerRef.CancelReservation(10)
}

func (r Reservation) CreateReservation(name, lastname string, dni int) {
	r.sellerRef.SetData(name, lastname, dni)
}
