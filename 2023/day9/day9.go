package main

import (
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

func part2() (int, error) {
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
		result += previousHistoryValue(history)
	}

	return result, nil
}

func previousHistoryValue(history []int) int {
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

	var previousValue = 0
	reversed := reverse(sequences)
	for i := 0; i < len(reversed)-1; i++ {
		b := reversed[i+1][0]
		previousValue = b - previousValue
	}
	return previousValue
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
	intArray := make([]int, len(arr))
	for i, str := range arr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArray[i] = num
	}

	return intArray
}
