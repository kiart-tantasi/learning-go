package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer
	fmt.Println("[Integer]")
	var _ int16 = 32_767
	var _ int32 = 2_147_483_647
	var _ int = 9_223_372_036_854_775_807
	var _ int = 9_223_372_036_854_775_807 // max is the same as int64

	// get max number of a type
	fmt.Printf("Max of int16 is %d\n", math.MaxInt16)
	fmt.Printf("Max of int32 is %d\n", math.MaxInt32)
	fmt.Printf("Max of int64 is %d\n\n", math.MaxInt64)

	// different ways to declare a variable
	var num1 int = 100
	var num2 = 100
	num3 := 100

	// declaring multiple variables
	fmt.Println("[Declaring Multiple Variables]")
	num4, num5, num6 := 100, 100, 100

	a, b, c := "string", 123, true
	fmt.Println(a, b, c)
	var d, e, f = "string", 123, true
	fmt.Println(d, e, f)
	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)
	// var d string, e int, f bool = "string", 123, true // this does not work

	// declare multiple variables with multiple types
	var (
		m string = "string"
		n int    = 123
		o bool   // to assign value later
	)
	o = true
	fmt.Println(m, n, o)

	// check type with "%T" and with "reflect"
	fmt.Println("\n[Checking Type]")
	var test1 int = 1
	var test2 string = "abc"
	var test3 bool = false
	var test4 int32 = 1
	// %T
	fmt.Printf("types are %T %T %T %T (%%T)\n", test1, test2, test3, test4)
	// reflect
	fmt.Printf("types are %s %s %s %s (reflect)\n", reflect.TypeOf(test1), reflect.TypeOf(test2), reflect.TypeOf(test3), reflect.TypeOf(test4))

	// to ignore usused error
	_ = num1 + num2 + num3 + num4 + num5 + num6
}
