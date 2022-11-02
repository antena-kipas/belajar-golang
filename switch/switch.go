package main
import "fmt"


func main () {
	name := "jangkrik"
	//switch biasa
	switch name {

	case "belalang" :
		fmt.Println("salah")
	case "semut" :
		fmt.Println("salah")
	case "jangkrik":
		fmt.Println("benar")
	default:
		fmt.Println("serangga apakah itu?")
	
	}


	//switch dengan short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("itu serangga berisik")
	case false:
		fmt.Println("serangga tidak berisik namun ada suaranya")
	}


	//switch tanpa komentar awal

	panjang := len(name)

	switch {
	case panjang > 10 :
		fmt.Println("nama lumayan panjang")
	case panjang > 5 :
		fmt.Println("nama cukup panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}