package main

import (
	"fmt"
)

func primeX(number int) int {
	count := 0 // Variabel untuk menghitung bilangan prima
	prime := 2 // Variabel untuk menyimpan bilangan prima

	for count < number {
		isPrime := true // Variabel untuk menandai apakah bilangan saat ini prima atau tidak

		// Memeriksa apakah bilangan saat ini prima
		for i := 2; i < prime; i++ {
			if prime%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++ // Jika bilangan saat ini prima, tambahkan ke hitungan
		}

		prime++ // Lanjut ke bilangan berikutnya
	}

	return prime - 1 // Mengembalikan bilangan prima terakhir yang ditemukan
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) //29
}
