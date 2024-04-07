package cmd



type resBuilder struct {
	vertical string
	rdate string
}

func (r *resBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v 
	return r
}


func (r *resBuilder) ReservationDate(date string) ReservationBuilder {
	r.rdate = date 
	return r
}

func (r *resBuilder) Build() Reservation {
	var builtReservation Reservation

	switch r.vertical {
		case "flight":
			builtReservation = flightReservation{reservationDate:r.rdate}
		case "hotel":
			builtReservation = hotelReservation{reservationDate:r.rdate}
	}

	return builtReservation
}