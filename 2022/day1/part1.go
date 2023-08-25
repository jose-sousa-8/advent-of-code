package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"errors"
)

func part1() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	currentCal := 0
	mostCaloriesElf := 0
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			if currentCal > mostCaloriesElf {
				mostCaloriesElf = currentCal
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

	return currentCal, nil
}
