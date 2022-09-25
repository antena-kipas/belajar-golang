package main

import "fmt"

func main() {
	var (
		name = "eko"
		e = name[0]
		eString = string(e)
	)

	fmt.Println(e)
	fmt.Println(eString)
}