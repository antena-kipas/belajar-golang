package main

import "fmt"

func main(){
	var person map[string]string = map[string]string{
		"name" : "janx",
		"alamat" : "Majalengka",		
	}

	person["judul"] = "belajar-golang"


	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"]) 

 	book := make(map[string]string)
	book["title"] = "belajar-golang"
	book["student"] = "ujanx"
	book["ups"] = "salah"
		
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
