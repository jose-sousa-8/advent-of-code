package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	scoreMap := map[string]int {
		"AX" : 4,
		"AY" : 8,
		"AZ" : 3,
		"BX" : 1,
		"BY" : 5,
		"BZ" : 9,
		"CX" : 7,
		"CY" : 2,
		"CZ" : 6,
	}

	score := 0
	for sc.Scan() {
		line := sc.Text()
		game := strings.Split(line, " ")
		res := game[0]+game[1]
		score += scoreMap[res]
	}

	return score, nil
}
