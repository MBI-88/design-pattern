package main

import "composite/cmd"

func main() {
	parend := cmd.NewComposite()

	parend.IsComposite()

	child := cmd.NewComposite()

	parend.AddChild(child)
	parend.IsComposite()
}