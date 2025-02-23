package main

import "fmt"

func main() {
	
	type NoKTP string

	var ktpYaya NoKTP = "11111111"

	fmt.Println(ktpYaya)

	var contoh string = "22222222"
	var convertContoh NoKTP = NoKTP(contoh);

	fmt.Println(contoh)
	fmt.Println(convertContoh)
}