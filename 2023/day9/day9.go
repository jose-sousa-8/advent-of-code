package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() (int, error) {
	buffer, _ := os.ReadFile("input")

	file := string(buffer)

	vals := strings.Split(file, "\r\n")
	var histories = [][]int{}
	for _, line := range vals {
		arr := convert(strings.Split(line, " "))
		histories = append(histories, arr)
	}

	var result = 0
	for _, history := range histories {
		result += nextHistoryValue(history)
	}

	return result, nil
}

func nextHistoryValue(history []int) int {
	sequences := [][]int{history}

	currentSequence := history
	for !allZero(currentSequence) {
		tmp := []int{}
		for i := 0; i < len(currentSequence)-1; i++ {
			tmp = append(tmp, int((currentSequence[i+1] - currentSequence[i])))
		}
		currentSequence = tmp
		sequences = append(sequences, tmp)
	}

	var nextVal = 0
	reversed := reverse(sequences)
	for i := 0; i < len(reversed)-1; i++ {
		b := reversed[i+1][len(reversed[i+1])-1]
		nextVal += b
	}
	return nextVal
}

func reverse(matrix [][]int) [][]int {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	return matrix
}

func allZero(sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			return false
		}
	}

	return true
}

func convert(arr []string) []int {
	// Convert strings to integers
	intArray := make([]int, len(arr))
	for i, str := range arr {
		// Parse the string to an integer
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting string '%s' to integer: %v\n", str, err)
		}
		intArray[i] = num
	}

	return intArray
}

// func part2() (string, error) {
// 	buffer, _ := os.ReadFile("1.txt")

// 	re := regexp.MustCompile(pattern)

// 	file := string(buffer)

// 	lines := strings.Split(file, "\r\n\r\n")
// 	steps := lines[0]
// 	nodes := map[string][]string{}
// 	for _, l := range strings.Split(lines[1], "\r\n") {
// 		matches := re.FindStringSubmatch(l)
// 		nodes[matches[1]] = []string{matches[2], matches[3]}
// 	}

// 	nodeSteps := []int{}
// 	startNodes := startNodes(nodes)
// 	for k, _ := range startNodes {
// 		var result = 0
// 		current := k
// 		for !strings.Contains(current, "Z") {
// 			for _, direction := range steps {
// 				result++
// 				if direction == 'L' {
// 					current = nodes[current][0]
// 				} else {
// 					current = nodes[current][1]
// 				}

// 				if strings.Contains(current, "Z") {
// 					nodeSteps = append(nodeSteps, result)
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return findLCM(nodeSteps).String(), nil
// }
