package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)


	var firstName = "Yaya Suryana"
	var y = firstName[0]
	var yString = string(y)

	fmt.Println(firstName)
	fmt.Println(y)
	fmt.Println(yString)
}