package cmd


type MediatorInt interface {
	SetColleagueC1(c1 *Colleague1)
	SetColleagueC2(c2 *Colleague2)
	SetState(s string)
}


func NewMediator() MediatorInt {
	return &mediator{}
}

