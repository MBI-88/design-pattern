package cmd




type flightInvoice struct {
	payment float64 
	date string
}

func (f flightInvoice) CalculateInvoce() float64 {
	return f.payment 
}

func (f *flightInvoice) SetInvoce(total float64, d string) {
	f.payment = total
	f.date = d
}

func (f flightInvoice) GetDate() string {
	return f.date
}