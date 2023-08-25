package main

import (
	"bufio"
	"os"
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
	lineNumber := 0
	elfGroup := [3]string {}
	for sc.Scan() {
		line := sc.Text()
		elfGroup[lineNumber] = line

		if (lineNumber == 2) {
			biggestIndex, midIndex, lowestIndex := min(len(elfGroup[0]), len(elfGroup[1]), len(elfGroup[2]))
			for	i := 0; i < len(elfGroup[lowestIndex]); i++ {
				c := elfGroup[lowestIndex][i]
				if strings.Contains(elfGroup[midIndex], string(c)) {
					
					if strings.Contains(elfGroup[biggestIndex], string(c)) {
						
						if (elfGroup[lowestIndex][i] >= 97) {
							score += int(elfGroup[lowestIndex][i]) - 96
							break
						}
						
						score += int(elfGroup[lowestIndex][i]) - 38
						break;
					}
				}
			}
			lineNumber = 0
			continue
		}

		lineNumber++
	}

	return score, nil
}

// lol
func min(x, y, z int) (int, int, int) {
	if x > y {
		if x > z {
			if (y > z) {
				return 0, 1, 2
			}
			return 0, 2, 1
		}
		return 2, 0, 1;
	}
	if z > y {
		return 2, 1, 0
	}
	if x > z {
		return 1, 0, 2
	}

	return 1, 2, 0
}
