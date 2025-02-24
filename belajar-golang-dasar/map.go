package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Yaya",
		"Address" : "Kuningan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["Address"])
	fmt.Println(person)

	book := make(map[string]string)

	book["title"] = "Ketika Cinta Bertasbih"
	book["author"] = "Yaya"
	book["wrong"]	= "salah"

	delete(book, "wrong")

	fmt.Println(book)
}