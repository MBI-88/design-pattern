package cmd

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
	SetReservationDate(string)
	SetPayment(float64)
}

type Invoice interface {
	CalculateInvoce() float64
	SetInvoce(float64, string)
	GetDate() string
}

type AbstractFactory interface {
	CreateReservation() Reservation
	CreateInvoice() Invoice
}

func NewAbstractFactory(v string) AbstractFactory {
	var factory AbstractFactory

	switch v {
	case "flight":
		factory = flighFactory{}

	case "hotel":
		factory = hoteFactory{}
	}

	return factory
}
