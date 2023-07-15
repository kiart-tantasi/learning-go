package main

import "fmt"

func main() {
	num := 3
	switch num {
	case 1:
		fmt.Print("value is 1")
	case 2:
		fmt.Print("value is 2")
	default:
		fmt.Printf("value %d does not exist in system\n", num)
	}
}

func something(a int) {

}
