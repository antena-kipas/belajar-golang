package main 

import "fmt"

func main() {
	var a = 10 
	var b = 20
	var c = a + b

	fmt.Println(c)
	
	var result = 10 + 100
	fmt.Println(result)

	a = a + 10
	fmt.Println(a)

	a += 10

	fmt.Println(a)


	a += 10
	a -= 10
	a *= 10
	a /= 10 
	a %= 10

	var i = 10
	i++
	fmt.Println(i)

	var negative = -100
	var positive = 100

	fmt.Println(negative)
	fmt.Println(positive)
}

