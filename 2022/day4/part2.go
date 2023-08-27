package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part2() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	score := 0
	for sc.Scan() {
		line := sc.Text()

		var pairs = strings.Split(line, ",")
		var pair1 = strings.Split(pairs[0], "-")
		var pair2 = strings.Split(pairs[1], "-")

		a1, _ := strconv.ParseInt(pair1[0], 10, 32)
		a2, _ := strconv.ParseInt(pair1[1], 10, 32)
		b1, _ := strconv.ParseInt(pair2[0], 10, 32)
		b2, _ := strconv.ParseInt(pair2[1], 10, 32)

		if between(a1, b1, b2) || between(a2, b1, b2) {
			score++
			continue
		}

		if between(b1, a1, a2) || between(b2, a1, a2) {
			score++
			continue
		}
	}

	return score, nil
}
