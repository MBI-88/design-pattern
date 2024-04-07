package main

import (
	"adaptor/cmd"
	"fmt"
)




func main() {
	adapter := cmd.NewAdapter()
	adapter.SetParams("John", "Do", 45652575)

	fmt.Println(adapter.GetArrayParams())
	fmt.Println(adapter.GetMapParams())
}