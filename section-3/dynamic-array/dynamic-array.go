package main

import "fmt"

func main() {
	names := []string{"Rohit", "Virat", "Rahul", "Bumrah"}   // this is a slice declaration dynamic length
	names2 := [4]string{"Rohit", "Virat", "Rahul", "Bumrah"} // this is an array declaration fixed length

	fmt.Printf("names %+v\n", names)
	fmt.Printf("names %+v\n", names[1:])
	fmt.Printf("names %+v\n", names[3:4])
	fmt.Printf("names2 %+v\n", names2[:3])

	items := make([]int, 3, 5)
	// 3 and 5 are the initial memory and cap of this dynamic array
	fmt.Printf("items %+v\n", items)
	fmt.Printf("items %+v , size: %d  , cap: %d\n", items, len(items), cap(items))
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	fmt.Printf("items %+v , size: %d  , cap: %d\n", items, len(items), cap(items))
}
