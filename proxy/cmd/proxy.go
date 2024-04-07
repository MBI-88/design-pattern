package cmd


type proxy struct {
	subject *hotelBoutique
}


func (p *proxy) Create(name, dni, card string) map[string]string {

	if p.subject == nil  {
		p.subject = new(hotelBoutique)
	}
	p.subject.CreateAccount(name, dni, card)
	return p.etJSON()
}


func (p proxy) etJSON() map[string]string {
	data := make(map[string]string)
	data["name"] = p.subject.name
	data["dni"] = p.subject.dni
	data["card"] = p.subject.card
	data["status"] = p.subject.status
	return data
}

func (p proxy) Cancel() map[string]string {
	p.subject.CancelAccount()
	return p.etJSON()
}