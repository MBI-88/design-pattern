package main

import (
	"chain/cmd"
)


func main() {
	ch := cmd.NewChain()
	ch.Handle("cmd") // output: No Receiver could handle the request!
	ch.Handle("pwd") // output: pwd  Receiver Finishing


}