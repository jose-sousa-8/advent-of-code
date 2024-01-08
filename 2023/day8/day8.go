package main

import (
	"math/big"
	"os"
	"regexp"
	"strings"
)

var pattern = `(\w+) = \(([^,]+), ([^)]+)\)`

type node struct {
	left     string
	right    string
	value    string
	stepsToZ int
}

func part1() (int, error) {
	buffer, _ := os.ReadFile("1.txt")

	re := regexp.MustCompile(pattern)

	file := string(buffer)

	lines := strings.Split(file, "\r\n\r\n")
	steps := lines[0]
	nodes := map[string][]string{}
	for _, l := range strings.Split(lines[1], "\r\n") {
		matches := re.FindStringSubmatch(l)
		nodes[matches[1]] = []string{matches[2], matches[3]}
	}

	current := "AAA"
	var result = 0
	for current != "ZZZ" {
		for _, direction := range steps {
			result++
			if direction == 'L' {
				current = nodes[current][0]
			} else {
				current = nodes[current][1]
			}

			if current == "ZZZ" {
				break
			}
		}
	}

	return result, nil
}

func part2() (string, error) {
	buffer, _ := os.ReadFile("1.txt")

	re := regexp.MustCompile(pattern)

	file := string(buffer)

	lines := strings.Split(file, "\r\n\r\n")
	steps := lines[0]
	nodes := map[string][]string{}
	for _, l := range strings.Split(lines[1], "\r\n") {
		matches := re.FindStringSubmatch(l)
		nodes[matches[1]] = []string{matches[2], matches[3]}
	}

	nodeSteps := []int{}
	startNodes := startNodes(nodes)
	for k, _ := range startNodes {
		var result = 0
		current := k
		for !strings.Contains(current, "Z") {
			for _, direction := range steps {
				result++
				if direction == 'L' {
					current = nodes[current][0]
				} else {
					current = nodes[current][1]
				}

				if strings.Contains(current, "Z") {
					nodeSteps = append(nodeSteps, result)
					break
				}
			}
		}
	}

	return findLCM(nodeSteps).String(), nil
}

func startNodes(nodes map[string][]string) map[string][]string {
	starters := map[string][]string{}
	for k, v := range nodes {
		if strings.Contains(k, "A") {
			starters[k] = v
		}
	}
	return starters
}

func findLCM(numbers []int) *big.Int {
	if len(numbers) == 0 {
		return big.NewInt(0)
	}

	result := big.NewInt(int64(numbers[0]))

	for i := 1; i < len(numbers); i++ {
		current := big.NewInt(int64(numbers[i]))
		result = lcm(result, current)
	}

	return result
}

func lcm(a, b *big.Int) *big.Int {
	product := new(big.Int)
	product.Mul(a, b)

	gcdAB := gcd(a, b)

	result := new(big.Int)
	result.Div(product, gcdAB)

	return result
}

func gcd(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		temp := new(big.Int)
		temp.Set(b)
		b.Mod(a, b)
		a.Set(temp)
	}
	return a
}
