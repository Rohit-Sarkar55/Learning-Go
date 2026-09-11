package main

import "fmt"

func modifyValue(val int) {
	val = val * 5
	fmt.Println("modifyValue:", val)
}

func modifyPointer(val *int) {
	*val = *val * 5
	fmt.Println("modifyPointer:", *val)
}

func main() {

	var a = 10
	var a_ptr = &a

	fmt.Println("a:", a, "address of a", &a)
	fmt.Printf("a ptr %+v\n", a_ptr)

	var b = 10
	modifyValue(b)
	fmt.Println("b:", b, "address of b", &b)

	modifyPointer(&b)
	fmt.Println("b:", b, "address of b", &b)
}
