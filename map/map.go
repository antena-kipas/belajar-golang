package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "janx",
		"alamat" : "Majalengka",		
	}

	person["judul"] = "belajar-golang"


	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"]) 
}
