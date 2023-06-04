package main

import "fmt"

func main() {
	myMap := map[string]int {
		"one": 1,
		"two": 2,
	}
		
	fmt.Println(myMap["one"])

	val, ok := myMap["ABC"]
	if ok == true {
		fmt.Println("value is", val)
	} else {
		fmt.Println("key does not exist in map")
	}
}
