package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part2() (string, error) {
	file, err := os.Open("input")
	if err != nil {
		return "", err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	nCrates := 8
	result := ""
	queues := make(map[int]*Queue)
	for i := 0; i < 9; i++ {
		queue := Queue{} // Initialize an empty queue
		queues[i] = &queue
		queues[i].printQueue()
	}

	for sc.Scan() && nCrates > 0 {
		line := sc.Text()
		crate := rune(line[1])
		if unicode.IsLetter(crate) {
			queues[0].Enqueue(string(crate))
		}
		for i := 1; i < 9; i++ {
			crate := rune(line[i*4+1])
			if unicode.IsLetter(crate) {
				queues[i].Enqueue(string(crate))
			}
		}
		nCrates--
	}

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		movements := strings.Split(line, " ")
		qtdT, _ := strconv.ParseInt(movements[1], 10, 8)
		qtd := int(qtdT)
		fromQueueT, _ := strconv.ParseInt(movements[3], 10, 8)
		fromQueue := int(fromQueueT) - 1
		toQueueT, _ := strconv.ParseInt(movements[5], 10, 8)
		toQueue := int(toQueueT) - 1

		crates := make([]string, qtd)
		for i := 0; i < qtd; i++ {
			crat, _ := queues[fromQueue].Dequeue()
			crates[i] = crat
		}

		queues[toQueue].PrependBlock(crates)
	}

	for i := 0; i < len(queues); i++ {
		queues[i].printQueue()
	}

	return result, nil
}
