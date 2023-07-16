package main

import "fmt"

func main() {
	m1 := []int{1, 2, 3}
	for index := range m1 {
		looked := make(map[int]bool)
		looked[index] = true
		iter(index, m1, make(map[int]bool))
	}
}

func iter(index int, m []int, looked map[int]bool) {
	_, ok := looked[index]
	if ok {
		return
	}

	fmt.Println(index)
	newLooked := copyMap(looked)
	newLooked[index] = true
	for i := range m {
		iter(i, m, newLooked)
	}
}

func copyMap(m map[int]bool) map[int]bool {
	newMap := make(map[int]bool)
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}

// 0
//   1 2
//   2 1

// 1
//   0 2
//   2 0

// 2
//   0 1
//   1 0
