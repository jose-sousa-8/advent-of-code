package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"errors"
)

func part2() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	currentCal := 0
	var mostCaloriesElfs = [] int {0, 0, 0}

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			for i, c := range mostCaloriesElfs {
				if currentCal > c {
					mostCaloriesElfs[i] = currentCal
					break
				}
			}
			currentCal = 0
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			return 0, errors.New("invalid calorie value")
		}

		currentCal += cal
	}

	return sum(mostCaloriesElfs), nil
}

func sum[T int](arr []T) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += int(arr[i])
	}

	return sum
}
