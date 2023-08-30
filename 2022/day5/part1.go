package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Queue []string

func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) Prepend(value string) {
	*q = append([]string{value}, *q...)
}

func (q *Queue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", fmt.Errorf("queue is empty")
	}

	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

func (q *Queue) printQueue() {
	for i := 0; i < len(*q); i++ {
		fmt.Printf("%v", (*q)[i])
	}
	fmt.Println()
}

func part1() (string, error) {
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

		for i := 0; i < qtd; i++ {
			crat, _ := queues[fromQueue].Dequeue()
			queues[toQueue].Prepend(crat)
		}
	}

	for i := 0; i < len(queues); i++ {
		queues[i].printQueue()
	}

	return result, nil
}
