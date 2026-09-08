package main

import "fmt"

func main() {
	var matrix [3][2]int

	fmt.Printf("Matrix length: %d and breadth : %d\n", len(matrix), len(matrix[0]))

	fmt.Printf("values %+v\n", matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = (i * len(matrix[0])) + j
		}
	}
	fmt.Printf("values %+v\n", matrix)
}
