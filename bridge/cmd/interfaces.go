package cmd

type SellerInt interface {
	CancelReservation(charge float64)
	SetData(name, lastname string, dni int)
}

type ReservationType interface {
	Cancel()
	CreateReservation(name, lastname string, dni int)
}

func NewReservation(reservationType string) ReservationType {
	var res ReservationType
	switch reservationType {
	case "normal":
		res = Reservation{&InstitutionSeller{}}
	case "premium":
		res = PremiumReservation{Reservation{&SmallScaleSeller{}}}
	}
	return res
}
