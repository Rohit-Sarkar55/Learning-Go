package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation      string
	InputA, InputB int
	Message        string
}

const division = "division"

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a: %d ", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b: %d ", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s", e.Operation,
		strings.Join(inputs, ","), e.Message)
}

func add(number ...int) int {
	defer fmt.Println("Sum finished")

	totl := 0
	for _, i := range number {
		total += i
	}
	return totl
}

func safeDivision(a int, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{Operation: "division", InputA: a,
			InputB: b, Message: "Division by Zero"}
	}
	return a / b, nil
}

func main() {
	fmt.Println(add(1, 2, 3))
	ans, err := safeDivision(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
	fmt.Println(safeDivision(40, 12))
}
