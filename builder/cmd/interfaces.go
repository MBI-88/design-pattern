package cmd



type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}


type ReservationBuilder interface {
	Vertical(string) ReservationBuilder
	ReservationDate(string) ReservationBuilder
	Build() Reservation
}


// Builder method
func NewReservationBuilder() ReservationBuilder {
	return &resBuilder{}
}