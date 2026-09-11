package main

import "fmt"

func main() {

	_map := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("map %+v  size : %d \n", _map, len(_map))

	//inorder to see type
	fmt.Printf("map is type of %T\n", _map)

	_map["a"] = 23
	fmt.Printf("map %+v  size : %d \n", _map, len(_map))

	var _map1 map[string]int // nil declaration

	map2 := make(map[string]int)

	if _map1 == nil {
		fmt.Printf("_map is nil\n")
	}

	if map2 == nil {
		fmt.Printf("map2 is nil\n")
	}

	value, found := _map["a"]
	if found {
		fmt.Printf("found value  %+v\n", value)
	}
	value2, found2 := _map["d"]
	if found2 {
		fmt.Printf("found value  %+v\n", value2)
	}
	// delete from maps

	delete(_map, "a")
	fmt.Printf("map %+v  size : %d \n", _map, len(_map))
}
