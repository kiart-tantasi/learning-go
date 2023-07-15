package main

import "fmt"

// Explanation
// while passing normal variables will create a new copy
// passing map will not create a new map and instead pass along the original map
// more info: https://forum.golangbridge.org/t/cant-index-pointer-to-map/26496/2

func main() {
	// edit without pointer
	m := make(map[int]int)
	sl := []int{}
	b := false
	i := 1
	s := SomeStruct{
		name: "John",
		age:  30,
	}
	editMap(m)
	editSlice(sl)
	editBool(b)
	editInt(i)
	editStruct(s)
	fmt.Println("[without pointer]")
	fmt.Println("edited map:", m)
	fmt.Println("edited slice:", sl)
	fmt.Println("edited bool:", b)
	fmt.Println("edited int:", i)
	fmt.Println("edited struct:", s)

	// edit with pointer
	m = make(map[int]int)
	sl = []int{}
	b = false
	i = 1
	s = SomeStruct{
		name: "John",
		age:  30,
	}
	editMapP(&m)
	editSliceP(&sl)
	editBoolP(&b)
	editIntP(&i)
	editStructP(&s)
	fmt.Println("\n[with pointer]")
	fmt.Println("edited map:", m)
	fmt.Println("edited slice:", sl)
	fmt.Println("edited bool:", b)
	fmt.Println("edited int:", i)
	fmt.Println("edited struct:", s)

}

// without pointer
func editMap(m map[int]int) {
	m[1] = 10
}
func editSlice(sl []int) {
	sl = append(sl, 10)
}
func editBool(b bool) {
	b = !(b)
}
func editInt(i int) {
	i = i + 1
}
func editStruct(s SomeStruct) {
	s.name = "new name"
	s.age = 123
}

// with pointer
func editMapP(m *map[int]int) {
	(*m)[1] = 10
}
func editSliceP(sl *[]int) {
	*sl = append(*sl, 10)
}
func editBoolP(b *bool) {
	*b = !(*b)
}
func editIntP(i *int) {
	*i = *i + 1
}
func editStructP(s *SomeStruct) {
	s.name = "new name"
	s.age = 123
}

type SomeStruct struct {
	name string
	age  int
}
