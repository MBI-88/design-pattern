package cmd

import "fmt"

type receiver struct {
}

func (r *receiver) Action(msg string) {
	fmt.Println(msg)
}