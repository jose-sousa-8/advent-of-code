package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	// miliseconds
	time int
	// milimeters
	distance int
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func trim(arr []string) []string {
	var res = []string{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != "" && arr[i] != " " {
			res = append(res, arr[i])
		}
	}

	return res
}

func convert(in []string) []int {
	numbers := []int{}
	for _, s := range in {
		n, err := strconv.Atoi(s)
		if err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func readRaces(scanner *bufio.Scanner) []race {
	times := []int{}
	distances := []int{}
	races := []race{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time:") {
			times = convert(trim(strings.Split(strings.Split(line, ":")[1], " ")))
			continue
		}

		distances = convert(trim(strings.Split(strings.Split(line, ":")[1], " ")))
		continue
	}

	for i := 0; i < len(times); i++ {
		races = append(races, race{time: times[i], distance: distances[i]})
	}

	return races
}

func getPossibleSolutions(race race) int {

	lowerLimit := 0
	upperLimit := 0

	for i := 1; i < race.time; i++ {
		// i 1, 2
		// moving time 60, 59
		// vel 1, 2
		// distance traveled = 1 * 60, 2 * 59
		movingTime := race.time - i
		velocity := i
		distanceTraveled := velocity * movingTime
		if distanceTraveled > race.distance {
			lowerLimit = i
			break
		}
	}

	for i := race.time - 1; i > 0; i-- {
		// i 60, 59
		// moving time 1, 2
		// vel 60, 59
		// distance traveled = 60 * 1, 59 * 2
		movingTime := race.time - i
		velocity := i
		distanceTraveled := velocity * movingTime
		if distanceTraveled > race.distance {
			upperLimit = i
			break
		}
	}

	return upperLimit - lowerLimit + 1
}

func multiplyArrayElements(arr []int) int {
	result := 1
	for _, value := range arr {
		result *= value
	}
	return result
}

func part1() (int, error) {
	file, err := os.Open("input1")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	var result = 1
	sc := bufio.NewScanner(file)
	races := readRaces(sc)
	fmt.Println(races)
	for _, race := range races {
		solutions := getPossibleSolutions(race)
		result *= solutions
	}

	return result, nil
}

func part2() (int, error) {

	file, err := os.Open("input1")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		// line := sc.Text()
		// fmt.Println(line)
	}

	return 0, nil
}
