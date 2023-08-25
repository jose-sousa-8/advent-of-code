package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part2() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	// x = lose, y = draw, z = win
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
	
	resMap := map[string]string {
		"AX" : "AZ",
		"AY" : "AX",
		"AZ" : "AY",
		"BX" : "BX",
		"BY" : "BY",
		"BZ" : "BZ",
		"CX" : "CY",
		"CY" : "CZ",
		"CZ" : "CX",
	}

	score := 0
	for sc.Scan() {
		line := sc.Text()
		game := strings.Split(line, " ")
		res := game[0]+game[1]
		score += scoreMap[resMap[res]]
	}

	return score, nil
}
