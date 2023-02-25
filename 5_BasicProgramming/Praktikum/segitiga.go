package main

import "fmt"

func main() {
	segitiga()
}

func segitiga() {
	var panjang int

	fmt.Println("Input Nilai: ")
	fmt.Scanf("%d", &panjang)

	for x := 1; x <= panjang; x++ {
		for y := panjang; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
