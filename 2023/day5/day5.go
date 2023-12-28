package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type _map struct {
	source      int
	destination int
	length      int
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func readMaps(input []string) [][]_map {
	var maps = make([][]_map, 7)
	maps[0] = append(maps[0], readMap(input, "seed-to-soil")...)
	maps[1] = append(maps[1], readMap(input, "soil-to-fertilizer")...)
	maps[2] = append(maps[2], readMap(input, "fertilizer-to-water")...)
	maps[3] = append(maps[3], readMap(input, "water-to-light")...)
	maps[4] = append(maps[4], readMap(input, "light-to-temperature")...)
	maps[5] = append(maps[5], readMap(input, "temperature-to-humidity")...)
	maps[6] = append(maps[6], readMap(input, "humidity-to-location")...)
	return maps
}

func isInMap(value int, m _map) (bool, int) {
	if value < m.source || value >= m.source+m.length {
		return false, -1
	}

	index := value - m.source
	return true, m.destination + index
}

func readMap(input []string, desc string) []_map {
	var m = []_map{}

	for i, l := range input {
		if strings.Contains(l, desc) {
			index := i + 1
			_l := input[index]
			for _l != "" && index < len(input)-1 {
				numbers := convert(strings.Split(_l, " "))
				m = append(m, _map{source: numbers[1], destination: numbers[0], length: numbers[2]})
				index++
				_l = input[index]
			}
			if _l != "" {
				numbers := convert(strings.Split(_l, " "))
				m = append(m, _map{source: numbers[1], destination: numbers[0], length: numbers[2]})
			}
			break
		}
	}
	return m
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

func readSeeds1(input []string) []int {
	var seeds []int
	for _, l := range input {
		if strings.Contains(l, "seeds:") {
			seedsStr := trim(strings.Split(strings.Split(l, ":")[1], " "))
			return convert(seedsStr)
		}
	}

	return seeds
}

func readSeeds2(input []string) []_map {
	var seeds []_map
	for _, l := range input {
		if strings.Contains(l, "seeds:") {
			seedsStr := trim(strings.Split(strings.Split(l, ":")[1], " "))
			for i := 0; i < len(seedsStr)-1; i++ {
				if i%2 == 0 {
					source, _ := strconv.Atoi(seedsStr[i])
					length, _ := strconv.Atoi(seedsStr[i+1])
					seeds = append(seeds, _map{source: source, length: length})
				}
			}
		}
	}

	return seeds
}

func trim(arr []string) []string {
	var res = []string{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != "" && arr[i] != " " {
			res = append(res, arr[i])
		}
	}

	return res
}

func part1() (int, error) {
	file, err := os.Open("input1")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	var input []string
	for sc.Scan() {
		line := sc.Text()
		input = append(input, line)
	}

	seeds := readSeeds1(input)
	maps := readMaps(input)

	result := MaxInt
	for _, seed := range seeds {
		next := seed
		for _, m := range maps {
			for _, _map := range m {
				found, destination := isInMap(next, _map)
				if found {
					next = destination
					break
				}
			}
		}

		if next < result {
			result = next
		}
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

	var input []string
	for sc.Scan() {
		line := sc.Text()
		input = append(input, line)
	}

	seeds := readSeeds2(input)
	maps := readMaps(input)
	result := MaxInt

	channel := make(chan int)

	for _, seed := range seeds {
		go func(seed _map, maps [][]_map, channel chan int) {
			fmt.Println("starting seed range ", seed)
			var r = MaxInt
			for i := seed.source; i < seed.source+seed.length; i++ {
				var next = i
				for j := 0; j < len(maps); j++ {
					for _, _map := range maps[j] {
						found, destination := isInMap(next, _map)
						if found {
							next = destination
							break
						}
					}
				}
				if next < r {
					r = next
				}
			}

			channel <- r
			fmt.Println("finished seed range ", seed, r)
		}(seed, maps, channel)
	}

	for i := 0; i < len(seeds); i++ {
		res := <-channel
		if res < result {
			result = res
		}
	}

	return result, nil
}
