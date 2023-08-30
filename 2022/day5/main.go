package main

import "fmt"

func main() {
	p1, err := part1()
	if err != nil {
		fmt.Println("error executing part 1")
	}
	fmt.Println(p1)
}
