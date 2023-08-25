package main

import (
	"bufio"
	"os"
	"strings"
)

func part1() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		return 0, err
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	score := 0
	for sc.Scan() {
		line := sc.Text()

		c1 := ""
		for	i := 0; i < len(line) / 2; i++ {
			c1 += string(line[i])	
		}

		c2 := ""
		for	i := len(line) / 2; i < len(line); i++ {
			c2 += string(line[i])
		}
		
		for	i := len(line) / 2; i < len(line); i++ {
			c2 += string(line[i])
		}
		
		for	i := 0; i < len(c1); i++ {
			if strings.Contains(c2, string(c1[i])){
				if (c1[i] >= 97) {
					score += int(c1[i]) - 96
					break
				}

				score += int(c1[i]) - 38
				break;
			}
		}
	}

	return score, nil
}
