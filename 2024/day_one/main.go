package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("part_one")
	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	arr1 := make([]int, 0, 1000)
	arr2 := make([]int, 0, 1000)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)
	result_part1 := 0
	result_part2 := 0
	var arr3 = make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		// could optimize since the arrays are sorted but w/e
		repeated := 0
		for j := 0; j < 1000; j++ {
			if arr2[j] < arr1[i] {
				continue
			}
			if arr2[j] == arr1[i] {
				repeated++
				continue
			}
			break
		}

		arr3 = append(arr3, arr1[i]*repeated)
		result_part1 += absDiff(arr1[i], arr2[i])
	}

	for i := 0; i < 1000; i++ {
		result_part2 += arr3[i]
	}

	fmt.Printf("The result for day 1 part 1 is: %d\n", result_part1)
	fmt.Printf("The result for day 1 part 2 is: %d\n", result_part2)
}
func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
