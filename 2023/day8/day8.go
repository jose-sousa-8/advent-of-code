package main

import (
	"os"
	"regexp"
	"strings"
)

var pattern = `(\w+) = \(([^,]+), ([^)]+)\)`

type node struct {
	left     string
	right    string
	value    string
	stepsToZ int
}

func part1() (int, error) {
	buffer, _ := os.ReadFile("1.txt")

	re := regexp.MustCompile(pattern)

	file := string(buffer)

	lines := strings.Split(file, "\r\n\r\n")
	steps := lines[0]
	nodes := map[string][]string{}
	for _, l := range strings.Split(lines[1], "\r\n") {
		matches := re.FindStringSubmatch(l)
		nodes[matches[1]] = []string{matches[2], matches[3]}
	}

	current := "AAA"
	var result = 0
	for current != "ZZZ" {
		for _, direction := range steps {
			result++
			if direction == 'L' {
				current = nodes[current][0]
			} else {
				current = nodes[current][1]
			}

			if current == "ZZZ" {
				break
			}
		}
	}

	return result, nil
}
