package main

import "fmt"

func main() {
	var i int
	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&i)

	if i%2 == 0 {
		fmt.Println("Bilangan Genap")
	} else {
		fmt.Println("Bilangan Ganjil")
	}
}
