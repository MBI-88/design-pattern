package cmd



type InterfaceX interface {
	IsComposite()
	AddChild(InterfaceX)
}




func NewComposite() InterfaceX {
	return &composite{}
} 