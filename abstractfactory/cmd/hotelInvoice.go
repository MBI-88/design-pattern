package cmd




type hotelInvoice struct{
	payment float64 
	date string
}


func (h hotelInvoice) CalculateInvoce() float64 {
	return h.payment 
}

func (h *hotelInvoice) SetInvoce(total float64, d string) {
	h.payment = total
	h.date = d
}

func (h hotelInvoice) GetDate() string {
	return h.date
}