package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Red   = "red"
	Green = "green"
	Blue  = "blue"
)

type set struct {
	greenCubes int
	redCubes   int
	blueCubes  int
}

type game struct {
	id   int
	sets []set
}

const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

func part1() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	result := 0

	for sc.Scan() {
		line := sc.Text()
		game := game{}
		readGame(line, &game)

		validGame := true
		for _, set := range game.sets {
			if set.redCubes > maxRedCubes ||
				set.greenCubes > maxGreenCubes ||
				set.blueCubes > maxBlueCubes {
				validGame = false
				break
			}
		}

		if validGame {
			result += game.id
		}
	}

	return result, nil
}

func readGame(line string, g *game) {
	_gameSets := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(strings.Split(_gameSets[0], " ")[1])
	g.id = gameId
	_sets := strings.Split(_gameSets[1], ";")
	for _, _set := range _sets {
		_cubes := strings.Split(_set, ",")
		red := 0
		blue := 0
		green := 0
		for _, _cube := range _cubes {
			c := strings.Split(_cube, " ")
			val, _ := strconv.Atoi(c[1])
			switch c[2] {
			case Blue:
				blue += val
			case Red:
				red += val
			case Green:
				green += val
			}
		}
		set := set{}
		set.redCubes = red
		set.greenCubes = green
		set.blueCubes = blue
		g.sets = append(g.sets, set)
	}
}

func part2() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	result := 0

	for sc.Scan() {
		line := sc.Text()
		game := game{}
		readGame(line, &game)
		minRed := -1
		minGreen := -1
		minBlue := -1
		for _, set := range game.sets {
			if set.redCubes > minRed {
				minRed = set.redCubes
			}
			if set.greenCubes > minGreen {
				minGreen = set.greenCubes
			}
			if set.blueCubes > minBlue {
				minBlue = set.blueCubes
			}
		}

		result += minRed * minGreen * minBlue
	}

	return result, nil
}
