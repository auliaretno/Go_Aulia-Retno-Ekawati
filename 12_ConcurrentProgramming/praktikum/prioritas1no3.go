package main

import (
	"fmt"
)

func main() {
	// buat buffered channel dengan kapasitas 10
	ch := make(chan int, 10)

	// jalankan goroutine untuk mengirim bilangan kelipatan 3 ke channel
	go func() {
		for i := 0; i <= 35; i++ {
			if i%3 == 0 {
				ch <- i // kirim bilangan ke channel
			}
		}
		close(ch) // menutup channel setelah selesai mengirim data
	}()

	// membaca data dari channel dan mencetaknya
	for n := range ch {
		fmt.Println(n)
	}
}
