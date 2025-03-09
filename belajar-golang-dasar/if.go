package main

import "fmt"

func main() {
	name := "Yaya Suryana"

	if name == "Yaya" {
		fmt.Println("Hallo Yaya")
	}else if name == "Amir Hamjah" {
		fmt.Println("Hallo Amir")
	}else{
		fmt.Println("Boleh Kenalan?")
	}


	// short statment di golang

	if banyak_karakter := len(name); banyak_karakter > 6{
		fmt.Println("Nama Terlalu Panjang")
	}else{
		fmt.Println("Nama Sudah Sesuai")
	}
}