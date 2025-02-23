package main

import "fmt"

func main() {
	const firstName string = "Yaya"
	const lastName = "Suryana"

	const (
		firstName2 = "Yaya"
		lastName2 = "Suryana"
	)

	fmt.Println(firstName+" "+lastName)

	fmt.Println(firstName2)
	fmt.Println(lastName2)
}