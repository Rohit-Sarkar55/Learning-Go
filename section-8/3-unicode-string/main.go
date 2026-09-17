package main

import "fmt"

func main() {
	a := "日本語"
	b := "abc"

	//three letter japanese word and three letter english string

	fmt.Printf("len a : %d\n", len(a))
	fmt.Printf("len b : %d\n", len(b))

	for _, ch := range a {
		fmt.Printf("%v\n", ch)
	}

	data := []rune{'日', '本', '語'}

	for _, ch := range data {
		fmt.Printf("datatype of rune characters: %t\n", ch)
	}
}
