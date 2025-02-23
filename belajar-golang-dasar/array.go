package main

import "fmt"

func main() {
	// set array secara terpisah
	var names [3]string

	names[0] = "Yaya"
	names[1] = "Suryana"
	names[2] = "Yaya Suryana"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	// mendefinisikan array secara langsung

	var values = [3]int {
		89,90,91,
	}

	fmt.Println(values)

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// function yang bisa digunakan di array

	var nilai = [...]int{100,101,103,104}
	fmt.Println(len(nilai))
	fmt.Println(nilai)

	// jika ingin set index tertentu di array
	nilai[0] = 10
	fmt.Println(nilai)


}