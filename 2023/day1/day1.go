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
			if isDigit(line[i]) {
				firstDigit = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if isDigit(line[i]) {
				lastDigit = line[i]
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

func part2() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	calibrationSum := 0

	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4
	numbers["five"] = 5
	numbers["six"] = 6
	numbers["seven"] = 7
	numbers["eight"] = 8
	numbers["nine"] = 9

	for sc.Scan() {
		line := sc.Text()
		var firstDigit int
		var lastDigit int
		// fmt.Print(line)
	first:
		for i := 0; i < len(line); i++ {
			char := line[i]
			if isDigit(char) {
				firstDigit, _ = strconv.Atoi(string(char))
				break
			}

			tmpStr := "" + string(char)
			for j := i + 1; j < len(line); j++ {
				char := line[i]
				if isDigit(char) {
					firstDigit, _ = strconv.Atoi(string(char))
					break first
				}
				tmpStr = tmpStr + string(line[j])
				_, ok := numbers[tmpStr]
				if ok {
					firstDigit = numbers[tmpStr]
					break first
				}
			}
		}

	second:
		for z := len(line) - 1; z >= 0; z-- {
			char := line[z]
			if isDigit(char) {
				lastDigit, _ = strconv.Atoi(string(char))
				break
			}

			tmpStr := "" + string(char)
			for y := z - 1; y > 0; y-- {
				char := line[z]
				if isDigit(char) {
					lastDigit, _ = strconv.Atoi(string(char))
					break
				}
				tmpStr = string(line[y]) + tmpStr
				_, ok := numbers[tmpStr]
				if ok {
					lastDigit = numbers[tmpStr]
					break second
				}
			}
		}

		// fmt.Printf(" => %v, %v\n", firstDigit, lastDigit)
		calibrationSum += (firstDigit*10 + lastDigit)
	}

	return calibrationSum, nil
}
