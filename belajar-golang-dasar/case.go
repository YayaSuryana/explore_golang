package main

import "fmt"

func main() {
	var name string

	name = "Yaya Suryana"

	switch name {
	case "Yaya":
		fmt.Println("Nama Tidak Sesuai")
	case "Yaya Suryana":
		fmt.Println("Nama Sesuai")
	default:
		fmt.Println("Boleh Kenalan?")
	}

	// Menggunakan sort statment
	switch length := len(name); length > 7 {
		case true:
			fmt.Println("Nama Panjang")
		case false:
			fmt.Println("Nama Pendek")
	}

	// membuat expression di case
	length := len(name)
	switch{
		case length > 10 :
			fmt.Println("Nama lengkap")
		case length > 5 :
			fmt.Println("Nama Panggilan")
		default:
			fmt.Println("Boleh Kenalan?")
	}
}