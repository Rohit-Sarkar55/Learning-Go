package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	text := "Hello World! Welcome to Go"

	regGo, err := regexp.Compile(`Go`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Text %s , matches Go: %t\n", text, regGo.MatchString(text))

	text2 := "Products Codes: P123, X342, P789"
	rProductP := regexp.MustCompile(`P\d+`)
	firstProduct := rProductP.FindString(text2)
	fmt.Println(firstProduct)

	allProducts := rProductP.FindAllString(text2, -1)
	fmt.Printf("All products: %v\n", allProducts)
}
