package main

import "fmt"

type Car struct {
    carType string
    fuelIn float64
}

func (car Car) menghitungJarak() float64 {
	return float64(car.fuelIn) / (1.5)
}

func main() {
	car := Car{
		carType: "Sedan",
		fuelIn:  12,
	}
	fmt.Println("Jarak tempuh mobil", car.carType, "adalah", car.menghitungJarak())
}