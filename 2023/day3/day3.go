package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const symbol = -1
const dot = -2
const gear = -3

type point struct {
	x int
	y int
}

type rowNumber struct {
	value  int
	points []point
}

func part1() (int, error) {

	sum := 0
	matrix := createMatrix1()

	for x := 0; x < len(matrix[0]); x++ {
		rowNumbers := readRowNumbers(matrix, x)
		for _, rowNumber := range rowNumbers {
			for _, indice := range rowNumber.points {
				adjacents := getAdjacents(matrix, indice.x, indice.y)
				if hasSymbol(adjacents) {
					sum += rowNumber.value
					break
				}
			}
		}
	}

	return sum, nil
}

func part2() (int, error) {

	sum := 0
	matrix := createMatrix2()

	for x := 0; x < len(matrix[0]); x++ {
		gears := readGears(matrix, x)
		fmt.Println(gears)
		for _, gear := range gears {
			adjancentParts := readAdjacentNumbers(matrix, gear)
			if len(adjancentParts) == 2 {
				sum += adjancentParts[0].value * adjancentParts[1].value
			}
		}
	}

	return sum, nil
}

func readAdjacentNumbers(matrix [][]int, gear point) []rowNumber {
	var adjacentNumbers []rowNumber
	var numbers []rowNumber

	numbers = readRowNumbers(matrix, gear.x)

	// read next line numbers
	if gear.x < len(matrix)-1 {
		numbers = append(numbers, readRowNumbers(matrix, gear.x+1)...)
	}

	// read previous next line numbers
	if gear.x > 0 {
		numbers = append(numbers, readRowNumbers(matrix, gear.x-1)...)
	}

	for _, number := range numbers {
		if numberIsAdjacentToGear(gear, number) {
			adjacentNumbers = append(adjacentNumbers, number)
		}
	}

	return adjacentNumbers
}

func numberIsAdjacentToGear(gear point, number rowNumber) bool {
	for _, point := range number.points {
		if arePointsAdjacent(gear, point) {
			return true
		}
	}

	return false
}

// we can assume these points are on the same row or the row above or bellow
func arePointsAdjacent(p1 point, p2 point) bool {
	if p1.y == p2.y ||
		p1.y == p2.y-1 ||
		p1.y == p2.y+1 {
		return true
	}

	return false
}

func readSameLineAdjacentNumbers(matrix [][]int, gear point) []rowNumber {
	var rowNumbers []rowNumber

	return rowNumbers
}

func readUpperLineAdjacentNumbers(matrix [][]int, gear point) []rowNumber {
	var rowNumbers []rowNumber

	return rowNumbers
}

func readLowerLineAdjacentNumbers(matrix [][]int, gear point) []rowNumber {
	var rowNumbers []rowNumber

	return rowNumbers
}

func readGears(matrix [][]int, row int) []point {
	var gears []point
	for i := 0; i < len(matrix); i++ {
		if matrix[row][i] == gear {
			gears = append(gears, point{x: row, y: i})
		}
	}

	return gears
}

func hasSymbol(s []int) bool {
	for _, a := range s {
		if a == symbol {
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
		if rowSlice[i] >= 0 {
			previous = append(previous, rowSlice[i])
			number.points = append(number.points, point{x: row, y: i})
			if i == len(rowSlice)-1 {
				val := previous[0]
				for i := 1; i < len(previous); i++ {
					val = val*10 + previous[i]
				}
				number.value = val
				rowNumbers = append(rowNumbers, number)
			}
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

func createMatrix1() [][]int {

	file, err := os.Open("input1")
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
			if unicode.IsDigit(char) {
				num, _ := strconv.Atoi(string(char))
				slice = append(slice, num)
				continue
			}
			if char == '.' {
				slice = append(slice, dot)
				continue
			}

			slice = append(slice, symbol)
		}

		matrix = append(matrix, slice)
	}

	return matrix
}

func createMatrix2() [][]int {

	file, err := os.Open("input2")
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
			if unicode.IsDigit(char) {
				num, _ := strconv.Atoi(string(char))
				slice = append(slice, num)
				continue
			}
			if char == '.' {
				slice = append(slice, dot)
				continue
			}
			if char == '*' {
				slice = append(slice, gear)
				continue
			}

			slice = append(slice, symbol)
		}

		matrix = append(matrix, slice)
	}

	return matrix
}

// could be greatly improved >D
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
