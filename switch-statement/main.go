package main

import "fmt"

func main() {
	something(2)
}

func something(a int) {
	switch a {
		case 1:
			fmt.Print("value is 1")
		case 2:
			fmt.Print("value is 2")
		default:
			fmt.Println("value with key", a, "does not exist")
	}
}
