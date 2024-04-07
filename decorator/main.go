package main

import (
	"decorator/cmd"
	"flag"
	"fmt"
	"os"
)

var (
	name *string
	age *int 
	gender *string
)

func init() {
	name = flag.String("name", "", "set user name")
	age = flag.Int("age", 0, "set user age")
	gender = flag.String("gender", "", "set user gender")

	flag.Usage = func() {
		info := fmt.Sprintf("[*] Casino ")
		info += "Varibles:\nname\nage\ngender\n"
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	casino := cmd.NewCasino()
	casino.Setcasino(*name, *gender, *age)

	for key, val := range casino.Getcasino() {
		fmt.Printf("%s:%s\n", key, val)
	}
}