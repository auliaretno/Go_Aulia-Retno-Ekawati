package main

import "fmt"

func Mapping(slice []string) map[string]int {
	result := make(map[string]int)

	for _, value := range slice {
		result[value]++
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}