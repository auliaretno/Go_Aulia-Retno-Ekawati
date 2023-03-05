package main

import (
    "fmt"
)

func generate(numRows int) [][]int {
    triangle := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        row := make([]int, i+1)
        row[0], row[i] = 1, 1
        for j := 1; j < i; j++ {
            row[j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
        triangle[i] = row
    }
    return triangle
}

func main() {
    fmt.Println(generate(5)) // Output: [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
