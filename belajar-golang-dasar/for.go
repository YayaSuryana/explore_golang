package main

import "fmt"

func main() {
	// for manual
	counter := 1

	for counter <= 10{
		fmt.Println("Perulangan Ke-", counter)
		counter++
	}

	fmt.Println("Perulangan Selesai")

	// for menggunakan init statment dan  post statment
	for counter := 1; counter <= 10; counter++{
		fmt.Println("Perulangan Ke-",counter)
	}

	// mengakses slice manual
	names := []string {"Yaya","Suryana","Yana","Surya"}
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	// mengkases slice dengan menggunakan range
	// for index, name := range names{
	// 	fmt.Println(index, name)
	// }
	// karna saya tidak akan memakai index jadi index bisa diganti dengan _
	for _, name := range names{
		fmt.Println(name)
	}
}