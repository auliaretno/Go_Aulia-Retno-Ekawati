package main

import "fmt"

func main() {
	var Nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&Nilai)

	if Nilai < 0 || Nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else if Nilai >= 80 {
		fmt.Println("A")
	} else if Nilai >= 65 {
		fmt.Println("B")
	} else if Nilai >= 50 {
		fmt.Println("C")
	} else if Nilai >= 35 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
