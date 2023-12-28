package main

import (
	"fmt"
)

func main() {
	p1, err := part1()
	if err != nil {
		fmt.Println("error executing part 1")
	}

	p2, err := part2()
	if err != nil {
		fmt.Println("error executing part 2")
	}

	fmt.Printf("Part 1 anwser is: %v\n", p1)
	fmt.Printf("Part 2 anwser is: %v", p2)
}
