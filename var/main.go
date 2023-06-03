package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num1 int = 100
	num2 := 100
	plus := num1 + num2
	fmt.Println(plus)

	num4, num5, num6 := 100, 100, 100
	plus = plus + num4 + num5 + num6
	fmt.Println(plus)

	var (
		num7 int
		num8 int
		num9 int
	)
	num7, num8, num9 = 100, 100, 100
	plus = plus + num7 + num8 + num9
	fmt.Println(plus)

	// multiple variables

	// work
	a, b, c := "string", 123, true
	fmt.Println(a, b, c)
	var d, e, f = "string", 123, true
	fmt.Println(d, e, f)
	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)

	var (
		m string = "string"
		n int    = 123
		o bool
	)
	o = true
	fmt.Println(m, n, o)

	// does not work
	// var x string, y int, z bool = "string", 123, true
	// fmt.Println(x, y, z)

	// declare without assigning
	var not1 string
	var not2 int
	var not3 bool
	fmt.Printf("value of not1, not2, not3: %s %d %t\n", not1, not2, not3)
	// print types
	fmt.Printf("types of not1, not2, not3 are %T %T %T\n", not1, not2, not3)

	// check type
	fmt.Println("check type of not1 by 'reflect':", reflect.TypeOf(not1))
}
