package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := strings.Clone(s1)
	//s3 := unsafe.StringData(s1)

	fmt.Println(s1 == s2)
	//fmt.Println(unsafe.StringData(s1) == s3)

	fmt.Println(strings.HasSuffix(s1, "cd"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))

	fruits := strings.SplitSeq("apple banana kiwi mango", " ")

	for fruit := range fruits {
		fmt.Println(fruit)
	}

}
