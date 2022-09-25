package main

import "fmt"

func main() {
	type Nama string 
	type Married bool

	var (
		name Nama = "tutu t0t0"
		marriedstatus Married = false
	)

	fmt.Println(name, marriedstatus)
}