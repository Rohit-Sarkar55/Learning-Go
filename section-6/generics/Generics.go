package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func Sum[T Number | string](items ...T) T {
	var total T
	for _, i := range items {
		total += i
	}
	return total
}

func main() {

	sum1 := Sum(1, 2, 3, 4)
	fmt.Printf("Type: %T,  sum1: %v\n", sum1, sum1)

	sum2 := Sum(11.32, 32.342, 341.231, 123.4668)
	fmt.Printf("Type: %T,  sum2: %v\n", sum2, sum2)

	sum3 := Sum(1.1, 2.1, 4.2, 1.4)
	fmt.Printf("Type: %T,  sum3: %v\n", sum3, sum3)

	sum4 := Sum("1.1", "Rohit", "sarkar", "1.4")
	fmt.Printf("Type: %T,  sum4: %v\n", sum4, sum4)
}
