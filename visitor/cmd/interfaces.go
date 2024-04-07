package cmd

type Node interface {
	Accept(Visitor)
	GetMessage() string
}

type Visitor interface {
	Visit(Node)
}

func NewNode(node string) Node {
	switch node {
	case "nodeB":
		return &nodeB{}
	default:
		return &nodeA{}
	}
}

func NewVistior() Visitor {
	return &visitor{}
}