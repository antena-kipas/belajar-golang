package main

import "fmt"

func main() {
	jarak_majalengka_kadipaten := 1
	//loop biasa
	for jarak_majalengka_kadipaten <= 10 {
		fmt.Println("loop ke ", jarak_majalengka_kadipaten)
		jarak_majalengka_kadipaten++
	}

	//loop dengan short statement

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)
	}

	//loop dengan range untuk slice dan map

	slice := []string{"ajax", "juki", "supri"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range mengambil semua data di dalam array,slice
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	//jika tidak ingin menggunakan i sebagai keterangan
	//dan ingin agar hanya valuenya saja maka i
	// digantikan dengan _
	// *bila ada variable yang tidak digunakanb maka golang akan error*

	for _, nilai := range slice {
		fmt.Println(nilai)
	}

	//for range untuk map

	person := make(map[string]string)

	person["nama"] = "adi"
	person["gelar"] = "sarjana"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
