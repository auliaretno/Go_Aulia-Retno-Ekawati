package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; i <= 7; i++ {
		fmt.Printf("%d ", i*x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go printMultiples(3)
	time.Sleep(25 * time.Second)
}
