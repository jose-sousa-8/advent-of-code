package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1() (float64, error) {
	file, err := os.Open("input1")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	result := float64(0)
	for sc.Scan() {
		line := sc.Text()
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := convert(strings.Split(numbers[0], " "))
		playerNumbers := convert(strings.Split(numbers[1], " "))
		var numberOfMatches = numberOfMatches(playerNumbers, winningNumbers)
		if numberOfMatches > 0 {
			result += math.Pow(2, float64(numberOfMatches-1))
		}
	}

	return result, nil
}

type gameSet struct {
	cards []card
}

type card struct {
	id             int
	winningNumbers []float64
	playerNumbers  []float64
	count          int
}

func part2() (int, error) {
	file, err := os.Open("input2")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	gameSet := gameSet{}
	for sc.Scan() {
		line := sc.Text()
		cardId, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := convert(strings.Split(numbers[0], " "))
		playerNumbers := convert(strings.Split(numbers[1], " "))

		card := card{id: cardId, winningNumbers: winningNumbers, playerNumbers: playerNumbers, count: 1}
		gameSet.cards = append(gameSet.cards, card)
	}

	for _, card := range gameSet.cards {
		var numberOfMatches = numberOfMatches(card.playerNumbers, card.winningNumbers)
		for i := card.id; i < card.id+numberOfMatches && i < len(gameSet.cards); i++ {
			gameSet.cards[i].count += card.count
		}
	}

	return numberOfCards(gameSet), nil
}

func numberOfCards(game gameSet) int {
	count := 0
	for _, card := range game.cards {
		count += card.count
	}
	return count
}

func numberOfMatches(playerNumbers []float64, winningNumbers []float64) int {
	count := 0
	for _, v := range playerNumbers {
		if contains(winningNumbers, v) {
			count++
		}
	}

	return count
}

func contains(arr []float64, v float64) bool {
	for _, val := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func convert(arr []string) []float64 {
	var res []float64
	for i := 0; i < len(arr); i++ {
		v, err := strconv.Atoi(arr[i])
		if err == nil {
			res = append(res, float64(v))
		}
	}

	return res
}
