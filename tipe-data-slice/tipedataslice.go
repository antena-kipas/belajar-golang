package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Febuary",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "diubah"
	//fmt.Println(slice1)

	//slice1[0] = "mei lagi"
	//fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "ozon")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)	
	fmt.Println(months)
	
	newSlice := make([]string, 2, 5,)
	newSlice[0] = "ozon"
	newSlice[1] = "droid"
    fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, 3, 6)
	copy(copySlice, newSlice)
	copySlice[2] = "statosfer"
	fmt.Println(copySlice)
	
}