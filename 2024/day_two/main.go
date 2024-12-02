package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1 is: %d\n", part1())
	fmt.Printf("part 2 is: %d\n", part2())
}
func part1() int {
	file, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	result_part1 := 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		valid, _ := processLevels(levels)
		if valid {
			result_part1++
		}
	}

	return result_part1
}

func part2() int {
	file, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	result_part2 := 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		valid, i := processLevels(levels)
		if valid {
			result_part2++
			continue
		}

		if i == 0 {
			// test by removing current and next
			levels_t1 := removeElement(levels, i)
			levels_t2 := removeElement(levels, i+1)
			valid1, _ := processLevels(levels_t1)
			valid2, _ := processLevels(levels_t2)
			if valid1 || valid2 {
				result_part2++
			}
		} else if i == len(levels)-1 {
			// test by removing previous and current
			levels_t1 := removeElement(levels, i-1)
			levels_t2 := removeElement(levels, i)
			valid1, _ := processLevels(levels_t1)
			valid2, _ := processLevels(levels_t2)
			if valid1 || valid2 {
				result_part2++
			}
		} else {
			// test by removing previous, current and next
			// test by removing previous and current
			levels_t1 := removeElement(levels, i-1)
			levels_t2 := removeElement(levels, i)
			levels_t3 := removeElement(levels, i+1)
			valid1, _ := processLevels(levels_t1)
			valid2, _ := processLevels(levels_t2)
			valid3, _ := processLevels(levels_t3)
			if valid1 || valid2 || valid3 {
				result_part2++
			}
		}
	}

	return result_part2
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func removeElement(slice []string, index int) []string {
	// Create a new slice without modifying the original
	newSlice := append([]string{}, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}

// if the levels are good returns true
// if not return the lower and higher index of problematic level
func processLevels(levels []string) (bool, int) {
	start, _ := strconv.Atoi(levels[0])
	prev, _ := strconv.Atoi(levels[1])
	increasing := false
	if absDiff(start, prev) > 3 {
		return false, 0
	} else if start == prev {
		return false, 0
	} else if start < prev {
		increasing = true
	}
	for i := 2; i < len(levels); i++ {
		n, _ := strconv.Atoi(levels[i])
		if prev == n {
			return false, i
		} else if increasing && prev > n {
			return false, i
		} else if !increasing && prev < n {
			return false, i
		} else if absDiff(prev, n) > 3 {
			return false, i
		}
		prev = n
	}
	return true, 0
}
