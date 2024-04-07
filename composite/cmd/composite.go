package cmd

import "fmt"



type composite struct {
	children []InterfaceX
}

func (c *composite) IsComposite() {
	if len(c.children) == 0 {
		fmt.Println("I'm a leaf")
		return
	}

	fmt.Println("I'm a composite")

	for _, child := range c.children {
		child.IsComposite()
	}

}

func (c *composite) AddChild(child InterfaceX) {
	c.children = append(c.children, child)
}

