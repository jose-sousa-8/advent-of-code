package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() int {
	file, err := os.Open("input")
	if err != nil {
		return 0
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		for i := 0; i < len(line)-4; i++ {
			arr := [4]byte{line[i], line[i+1], line[i+2], line[i+3]}
			fmt.Printf("array %v\n", arr)
			if hasDuplicate(arr) {
				fmt.Printf("duplicate %v\n", arr)
				continue
			}

			fmt.Printf("Not duplicate %v\n", arr)
			return i + 4
		}
	}
	return 0
}

func hasDuplicate(arr [4]byte) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if j != i && arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}
