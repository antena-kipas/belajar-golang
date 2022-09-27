package main 

import "fmt"

func main() {
	var (
		name1 = "Tutu"
		name2 = "Toto"
		name3 = "janx"
		name4 = "dudu"

		result bool = name1 == name2 
		//pasti hasilnya false

		hasil2 bool = name3 > name4
		//pasti hasilnya True 
	)

	fmt.Println(result)
	fmt.Println(hasil2)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}