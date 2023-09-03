package main

import (
	"bufio"
	"os"
)

func part2() int {
	file, err := os.Open("input")
	if err != nil {
		return 0
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		for i := 0; i < len(line)-14; i++ {
			arr := make([]byte, 14)
			counter := 0
			for j := i; j < i+14; j++ {
				arr[counter] = line[j]
				counter++
			}

			if hasDuplicate(arr) {
				continue
			}

			return i + 14
		}
	}
	return 0
}
