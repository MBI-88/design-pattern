package cmd



type hoteFactory struct {}

func (hf hoteFactory) CreateReservation() Reservation {
	return new (hotelReservation)
}

func (hf hoteFactory) CreateInvoice() Invoice {
	return new (hotelInvoice)
}


type flighFactory struct{}

func (ff flighFactory) CreateReservation() Reservation{
	return new (flightReservation)
}

func (ff flighFactory) CreateInvoice() Invoice {
	return new (flightInvoice)
}