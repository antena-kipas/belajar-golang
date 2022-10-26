package main 
import "fmt"

func main () {
	var nama = "ante"
	var length = len(nama)

	if length > 5 {
		fmt.Println("nama anda sudah benar")
	} else {
		fmt.Println("nama anda salah")
	}

	fmt.Println(length)

	var kesan = "seru"

	if huruf := len(kesan) ; huruf > 5 {
		fmt.Println("why?")
	} else {
		fmt.Println("terima kasih")
	}


}	

