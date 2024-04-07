package cmd


type StateInt interface {
	Op1(*context)
	Op2(*context)
}



func NewContextA() StateInt {
	return &contextA{}
}

func NewContex() *context {
	c := new(context)
	c.setState(NewContextA())
	return c
}