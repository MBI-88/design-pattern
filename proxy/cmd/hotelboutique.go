package cmd


type hotelBoutique struct {
	name string
	dni string
	card string
	status string

}

func (h *hotelBoutique) AccountInfo() hotelBoutique {
	return *h
}

func (h *hotelBoutique) CreateAccount(name, dni, card string) {
	h.card = card
	h.dni = dni
	h.name = name
	h.status = "Created"
}

func (h *hotelBoutique) CancelAccount()  {
	h.status = "Canceled"
}