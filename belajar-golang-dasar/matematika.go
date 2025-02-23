package main

import "fmt"

func main() {
	a := 5
	b := 5
	c := a + b
	fmt.Println(c)

	// augmented assignment

	var i = 10
	i += 10

	fmt.Println(i)

	i += 5

	fmt.Println(i)


	// unary operator

	var j = 1

	j++
	fmt.Println(j)

	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
	j-- 
	fmt.Println(j)
}