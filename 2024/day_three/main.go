package main

import (
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed input
var text string

func main() {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)
	result := 0
	for _, match := range matches {
		var x, y int
		fmt.Sscanf(match, "mul(%d,%d)", &x, &y)
		result += x * y
	}
	fmt.Println(result)
}
