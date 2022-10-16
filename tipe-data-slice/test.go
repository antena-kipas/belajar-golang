package main

import "fmt"

func main() {
	var months = [...]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
	}

	var slice =  months[10:]
	fmt.Println(slice)

	var slice1 = append(slice, "ozon")
	fmt.Println(slice1)
	slice1[0] = "bukan Desember"
	fmt.Println(slice1)


}