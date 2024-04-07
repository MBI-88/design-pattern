package main

import (
	"fmt"
	"proxy/cmd"
)


func printJson(m map[string]string) {
	fmt.Println("{")
	for key, value := range m {
		fmt.Printf("  %s: %s,\n",key,value)
	}
	fmt.Println("}")
}

func main() {

	api := cmd.NewProxy()

	resp := api.Create("John", "12345678V", "2345987A45")

	printJson(resp)

	resp = api.Cancel()

	printJson(resp)


}