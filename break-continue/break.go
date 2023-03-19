package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("contoh ke", i)
	}

	for o := 0; 0 < 10; o++ {
		fmt.Println("contoh2 ke", o)

		if o == 5 {
			break
		}
	}
}
