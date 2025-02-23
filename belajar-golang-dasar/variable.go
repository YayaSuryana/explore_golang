package main

import "fmt"

func main() {
	// cara 1
	var name string
	name = "Yaya"

	fmt.Println(name)

	// cara kedua
	surename :="Yaya Suryana"
	fmt.Println(surename)

	// cara ketiga
	var (
		firstName = "Yaya"
		lastName	= "Suryana"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}