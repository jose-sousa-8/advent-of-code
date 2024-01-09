package main

import (
	"fmt"
	"os"
	"strings"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	x     int
	y     int
}

type BinaryTree struct {
	root *BinaryNode
}

func part1() (int, error) {
	buffer, _ := os.ReadFile("input")

	file := string(buffer)

	m := [][]rune{}
	for _, line := range strings.Split(file, "\r\n") {
		m = append(m, []rune(line))
	}

	node := getNode('S', m)
	fmt.Println(node)

	return 0, nil
}

func getNode(r rune, m [][]rune) *BinaryNode {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == r {
				return &BinaryNode{left: nil, right: nil, x: i, y: j}

			}
		}
	}

	return nil
}

func part2() (int, error) {
	buffer, _ := os.ReadFile("input")

	file := string(buffer)

	for _, line := range strings.Split(file, "\r\n") {
		fmt.Println(line)
	}

	return 0, nil
}
