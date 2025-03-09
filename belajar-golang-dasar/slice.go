package main

import "fmt"

func main() {
	
	var names = [...]string {"Yaya", "Suryana", "Raka", "Sutarman", "Jaka", "Wijaya"}

	// Membuat slice dari array yang sudah ada
	// ditentukan batas bawah dan akhirnya
	slice1 := names[2:5]
	fmt.Println(slice1)

	// mulai dari index yang di tentukan ke index terakhir
	// cara penulisan slice yang manual
	var slice2 []string = names[2:]
	fmt.Println(slice2)

	// mulai dari index awal / 0 hingga yang di tentukan
	slice3 := names[:4]
	fmt.Println(slice3)

	// mengambil seluruh data array
	var slice4 []string = names[:]
	fmt.Println(slice4)

	// function di slice
	var days = [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	fmt.Println(daySlice1)

	// jika menambahkan data baru ke array dengan append maka slice akan membuat array baru dan terpisah dengan array sebelumnya
	daySlice2 := append(daySlice1, "Libur Tambahan")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice2)
	// days dan dayslice1 di awal tidak terpengaruh oleh dayslice2 karna terpisah atau slice sudah membuat array baru
	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	// membuat array dari awal dengan slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Yaya"
	newSlice[1] = "Yaya"
	// newSlice[2] = "Yaya" /error karna untuk menambah array harus dengan append

	newSlice1 := append(newSlice, "Yaya", "Yaya","Yaya","Yaya")
	fmt.Println(newSlice)
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// akan terlihat perbedaan dimana dengan memakai make akan membuat array dari awal dengan slice tidak memanfaatkan array yang sudah ada
	// dan dengan membedakan antara length dengan capacity bertujuan agar jika ditambahkan data baru tidak langsung membuat aray baru melainkan akan dibuat array baru ketika capacity penuh
	
	// copy slice
	fromSlice 	:= days[:]
	toSlice		:= make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	
	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	// bedakan antara array dan slice
	iniArray := [...]int {1,2,3,4,5}
	iniSlice := []int {1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	

}