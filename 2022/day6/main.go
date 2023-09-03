package main

import "fmt"

func main() {
	p1 := part1()
	fmt.Printf("part1 %v\n", p1)

	p2 := part2()
	fmt.Printf("part2 %v", p2)
}

func hasDuplicate[T comparable](arr []T) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if j != i && arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}
