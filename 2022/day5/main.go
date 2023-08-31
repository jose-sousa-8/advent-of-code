package main

import "fmt"

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

func (q *Queue) PrependBlock(values []string) {
	for i := len(values) - 1; i >= 0; i-- {
		q.Prepend(values[i])
	}
}

func main() {
	p1, err := part1()
	if err != nil {
		fmt.Println("error executing part 1")
	}
	fmt.Println("part1 " + p1)

	p2, err := part2()
	if err != nil {
		fmt.Println("error executing part 1")
	}
	fmt.Println("part2 " + p2)
}
