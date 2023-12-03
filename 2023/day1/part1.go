package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	calibrationSum := 0

	for sc.Scan() {
		line := sc.Text()
		var firstDigit byte
		var lastDigit byte
		for i := 0; i < len(line); i++ {
			char := line[i]
			if isDigit(char) {
				firstDigit = char
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if isDigit(char) {
				lastDigit = char
				break
			}
		}

		// Concatenate characters as a string
		str := string(firstDigit) + string(lastDigit)

		// Convert the string to an integer
		result, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
		}
		calibrationSum += result
	}

	return calibrationSum, nil
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
