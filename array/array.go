package main 
import "fmt"

func main(){
	var nama[3]string
	nama[0] = "OZON"
	nama[1] = "asep garpu"
	nama[2] = "Opet"

	fmt.Println(nama)
	fmt.Println(nama[2])
	
	var angka = [3]int {
	    6,
	    7,
	    8,
	}
	
	fmt.Println(angka)
	fmt.Println(angka[1])
	angka[0] = 11
	fmt.Println(angka[0])
}
