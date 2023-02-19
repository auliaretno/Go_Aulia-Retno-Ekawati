package main

import "fmt"

func main() {
	var sisiA, sisiB, tinggi int

	// Input sisiA, sisiB dan tinggi
	fmt.Print("Masukkan sisi A: ")
	fmt.Scanln(&sisiA)
	fmt.Print("Masukkan sisi B: ")
	fmt.Scanln(&sisiB)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	// Luas Trapesium
	L := (sisiA * sisiB * tinggi) / 2
	fmt.Println(L)
}
