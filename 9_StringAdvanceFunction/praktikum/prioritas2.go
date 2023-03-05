package main

import "fmt"

func caesarOffset(input string, offset int) string {
	result := []rune{}

	for _, char := range input {
		char = 'a' + (char-'a'+rune(offset))%26
		result = append(result, char)
	}
	return string(result)
}

func main() {
	fmt.Println(caesarOffset("abc", 3))
	fmt.Println(caesarOffset("alta", 2))
	fmt.Println(caesarOffset("alterraacademy", 10))
	fmt.Println(caesarOffset("abcdefghijklmnopqrstuvwxyz", 1))
	fmt.Println(caesarOffset("abcdefghijklmnopqrstuvwxyz", 1000))

}
