package main

import "fmt"

// =================================
type Counter struct {
	value int
}

type CounterI interface {
	Increment()
	Value() int
}

// =================================

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func updateCounterInterface(c CounterI) {
	c.Increment()
}

func updateCounterNonInterface(c Counter) {
	c.Increment()
}

func newCounterInterface(value int) CounterI {
	c := Counter{value: value}
	return &c
}

func newCounterNonInterface(value int) Counter {
	c := Counter{value: value}
	return c
}

// =================================

func main() {
	// [INTERFACE]
	// create an interface var
	c1 := newCounterInterface(10)
	fmt.Printf("[INTERFACE] Value before updating is %d\n", c1.Value())
	// Pass a a non-pointer var
	updateCounterInterface(c1)
	fmt.Printf("[INTERFACE] Value after updating is %d\n", c1.Value())

	// [STRUCT]
	c2 := newCounterNonInterface(10)
	fmt.Printf("[STRUCT] Value before updating is %d\n", c2.Value())
	// Pass a a non-pointer var
	updateCounterNonInterface(c2)
	fmt.Printf("[STRUCT] Value after updating is %d\n", c2.Value())

	// [Summary]
	// Passing an interface var as a non-pointer var will also modify the original struct while passing a struct var won't do so.
}
