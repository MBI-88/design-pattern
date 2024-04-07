package cmd


// Base Interface
type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}

// Flght Interface
type FlightReservation interface {
	Reservation
	AddExtraLuggageAllowance(peices int)
}

// Hotel Interface
type HotelReservation interface {
	Reservation
	ChangeType(name string)
}

// Trip interface
type Trip interface {
	AddReservation(r Reservation)
	CalculateCancellationFee() float64
}

// Factory method
func NewReservationFactory(vertical, reservationDate string ) Reservation {
	switch vertical {
	case "hotel":
		return &hotelReservation{reservationDate:reservationDate}

	case "flight":
		return &flightReservation{reservationDate:reservationDate}

	default:
		return nil
	}

}

// Factory method
func NewTrip() Trip {
	return &trip{}
}