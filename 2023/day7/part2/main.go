package main

import (
	"fmt"
)

func main() {
	p2, err := part2()
	if err != nil {
		fmt.Println("error executing part 2")
	}

	fmt.Printf("Part 2 anwser is: %v", p2)
}
