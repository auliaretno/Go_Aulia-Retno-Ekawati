package main

import "fmt"

func main() {
	var Bilangan int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&Bilangan)

	fmt.Printf("Faktor dari %d adalah: ", Bilangan)
	for i := 1; i <= Bilangan; i++ {
		if Bilangan%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
