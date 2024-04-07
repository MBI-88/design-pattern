package cmd


type context struct {
	state StateInt
	chacked bool
	successful bool
	name string
	payed bool
}

func (a *context) Op1() {
	a.chacked = true
	a.successful = true
	a.name = "Status A"
	a.payed = true
	a.state.Op1(a)
}

func (a *context) Op2() {
	a.chacked = true
	a.successful = false
	a.name = "Status B"
	a.payed = false
	a.state.Op2(a)
}

func (a *context) setState(s StateInt) {
	a.state = s
}