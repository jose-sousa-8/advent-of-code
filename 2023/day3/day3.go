package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

type rowNumber struct {
	value   int
	indices []point
}

func part1() (int, error) {

	sum := 0
	matrix := createMatrix()

	for x := 0; x < len(matrix[0]); x++ {
		rowNumbers := readRowNumbers(matrix, x)
		for _, rowNumber := range rowNumbers {
			for _, indice := range rowNumber.indices {
				adjacents := getAdjacents(matrix, indice.x, indice.y)
				if contains(adjacents, -2) {
					sum += rowNumber.value
					break
				}
			}
		}
	}

	return sum, nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readRowNumbers(matrix [][]int, row int) []rowNumber {

	var rowSlice = matrix[row]
	rowNumbers := []rowNumber{}

	previous := []int{}
	number := rowNumber{}
	for i := 0; i < len(rowSlice); i++ {
		if rowSlice[i] > 0 {
			previous = append(previous, rowSlice[i])
			number.indices = append(number.indices, point{x: row, y: i})
			continue
		}

		if len(previous) > 0 {
			val := previous[0]
			for i := 1; i < len(previous); i++ {
				val = val*10 + previous[i]
			}
			number.value = val
			rowNumbers = append(rowNumbers, number)
		}

		previous = []int{}
		number = rowNumber{}
	}

	return rowNumbers
}

func createMatrix() [][]int {

	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	matrix := [][]int{}
	for sc.Scan() {
		line := sc.Text()
		slice := []int{}
		for _, char := range line {
			if isDigit(char) {
				num, _ := strconv.Atoi(string(char))
				slice = append(slice, num)
				continue
			}
			if char == '.' {
				slice = append(slice, -1)
				continue
			}

			slice = append(slice, -2)
		}

		matrix = append(matrix, slice)
	}

	return matrix
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func getAdjacents(matrix [][]int, x int, y int) []int {
	result := []int{}
	if x == 0 {
		if y == 0 {
			result = append(result, matrix[x][y+1])
			result = append(result, matrix[x+1][y])
			result = append(result, matrix[x+1][y+1])
			return result
		}
		if y == len(matrix[x])-1 {
			result = append(result, matrix[x][y-1])
			result = append(result, matrix[x+1][y-1])
			result = append(result, matrix[x+1][y])
			return result
		}

		result = append(result, matrix[x][y-1])
		result = append(result, matrix[x][y+1])
		result = append(result, matrix[x+1][y-1])
		result = append(result, matrix[x+1][y+1])
		result = append(result, matrix[x+1][y])
		return result
	}

	if x == len(matrix)-1 {
		if y == 0 {
			result = append(result, matrix[x][y+1])
			result = append(result, matrix[x-1][y])
			result = append(result, matrix[x-1][y+1])
			return result
		}
		if y == len(matrix[x])-1 {
			result = append(result, matrix[x][y-1])
			result = append(result, matrix[x-1][y-1])
			result = append(result, matrix[x-1][y])
			return result
		}

		result = append(result, matrix[x][y-1])
		result = append(result, matrix[x][y+1])

		result = append(result, matrix[x-1][y])
		result = append(result, matrix[x-1][y-1])
		result = append(result, matrix[x-1][y+1])
		return result
	}

	if y == 0 {
		result = append(result, matrix[x+1][y])
		result = append(result, matrix[x-1][y])

		result = append(result, matrix[x-1][y+1])
		result = append(result, matrix[x][y+1])
		result = append(result, matrix[x+1][y+1])
		return result
	}

	if y == len(matrix[x])-1 {
		result = append(result, matrix[x+1][y])
		result = append(result, matrix[x-1][y])

		result = append(result, matrix[x-1][y-1])
		result = append(result, matrix[x][y-1])
		result = append(result, matrix[x+1][y-1])
		return result
	}

	result = append(result, matrix[x][y-1])
	result = append(result, matrix[x][y+1])

	result = append(result, matrix[x-1][y])
	result = append(result, matrix[x-1][y+1])
	result = append(result, matrix[x-1][y-1])

	result = append(result, matrix[x+1][y])
	result = append(result, matrix[x+1][y-1])
	result = append(result, matrix[x+1][y+1])

	return result
}

// func part2() (int, error) {
// 	file, err := os.Open("input")
// 	if err != nil {
// 		fmt.Printf("Error reading file %v", file.Name())
// 	}

// 	defer file.Close()

// 	sc := bufio.NewScanner(file)
// 	result := 0

// 	for sc.Scan() {
// 		line := sc.Text()
// 	}

// 	return result, nil
// }
