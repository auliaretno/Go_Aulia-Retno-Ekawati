package main

import (
	"fmt"
)

func main() {
	// buat channel 
	ch := make(chan int)

	// jalankan goroutine untuk mengirim bil kelipatan 3 ke channel
	go func() {
		for i := 0; i <= 25; i++ {
			if i%3 == 0 {
				ch <- i // mengirim bil ke channel
			}
		}
		close(ch) // menutup channel setelah selesai mengirim data
	}()

	// membaca data dari channel dan mencetaknya
	for n := range ch {
		fmt.Println(n)
	}
}