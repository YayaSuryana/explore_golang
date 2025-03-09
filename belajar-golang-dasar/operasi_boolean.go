package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsensi = 88

	var lulusNilaiAkhir bool = nilaiAkhir >90
	var lulusAbsensi bool = nilaiAbsensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi 

	fmt.Println(lulus)
}