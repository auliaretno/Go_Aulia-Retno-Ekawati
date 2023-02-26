package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	var sumLR int
	for i := 0; i < len(matrix); i++ {
		sumLR += matrix[i][i]
	}

	var sumRL int
	for i := 0; i < len(matrix); i++ {
		sumRL += matrix[i][len(matrix)-1-i]
	}

	diff := sumLR - sumRL
	if diff < 0 {
		diff = -diff
	}

	fmt.Printf("Matriks: %v\n", matrix)
	fmt.Printf("Selisih absolut antara jumlah diagonalnya: %d\n", diff)
}