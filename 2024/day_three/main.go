package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
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

	fmt.Println(part2())
}

func part2() int {
	pattern := `(do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\))`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(text, -1)

	mulEnabled := true
	result := 0
	for _, match := range matches {
		if match == "do()" {
			mulEnabled = true
		} else if match == "don't()" {
			mulEnabled = false
		} else if strings.HasPrefix(match, "mul(") && mulEnabled {
			var x, y int
			fmt.Sscanf(match, "mul(%d,%d)", &x, &y)
			result += x * y
		}
	}

	return result
}
