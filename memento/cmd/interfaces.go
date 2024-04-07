package cmd


type OriginatorInt interface {
	GetState() string
	SetState(state string)
	GetMemento() memento
	Restore(me memento)
}


func Creataker() OriginatorInt {
	return &originator{}
}
