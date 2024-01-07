package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func hasSmallerHighCard(a string, b string) bool {
	for i := 0; i < len(a); i++ {
		if cards[a[i]] == cards[b[i]] {
			continue
		}

		return cards[a[i]] < cards[b[i]]
	}

	return false
}

func (a hands) Len() int { return len(a) }
func (a hands) Less(i, j int) bool {
	return a[i].pokerHand < a[j].pokerHand || (a[i].pokerHand == a[j].pokerHand && hasSmallerHighCard(a[i].cards, a[j].cards))
}
func (a hands) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

var cards map[byte]Card

type PokerHand int

type Card int

const (
	Joker = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Queen
	King
	Ace
)

const (
	HighCard PokerHand = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hands []hand

type hand struct {
	cards     string
	bid       int
	pokerHand PokerHand
}

func readPokerHand(handStr string) PokerHand {
	labels := map[rune]int{}

	for _, c := range handStr {
		_, exists := labels[c]
		if exists {
			labels[c]++
			continue
		}

		labels[c] = 1
	}

	pokerHand := labelsToPokerHand(labels)

	if labels['J'] < 1 {
		return pokerHand
	}

	return improvePokerHand(labels, labels['J'], pokerHand)
}

// we can always assume there is at least one joker
func improvePokerHand(labels map[rune]int, nJokers int, pokerHand PokerHand) PokerHand {

	// means five of a kind , four of a kinda and full house can all be converted to five of a kind
	if len(labels) <= 2 {
		return FiveOfAKind
	}

	// JJJ12 || JJ112 | J1122 | J1112
	if len(labels) == 3 {
		if nJokers >= 2 {
			return FourOfAKind
		}
		if nJokers == 2 {
			return FullHouse
		}

		// nJokers == 1
		if pokerHand == ThreeOfAKind {
			return FourOfAKind
		}

		return FullHouse
	}

	if len(labels) == 4 {
		return ThreeOfAKind
	}

	return OnePair
}

func labelsToPokerHand(labels map[rune]int) PokerHand {
	if len(labels) == 1 {
		return FiveOfAKind
	}
	if len(labels) == 2 {
		if isFullHouse(labels) {
			return FullHouse
		}
		return FourOfAKind
	}
	if len(labels) == 5 {
		return HighCard
	}
	if len(labels) == 3 {
		if isThreeOfAkind(labels) {
			return ThreeOfAKind
		}
		return TwoPair
	}

	return OnePair
}

func isThreeOfAkind(labels map[rune]int) bool {
	if len(labels) != 3 {
		return false
	}

	if containsValue(labels, 3) {
		return true
	}

	return false
}

func containsValue(labels map[rune]int, value int) bool {
	for _, v := range labels {
		if v == value {
			return true
		}
	}
	return false
}

func isFullHouse(labels map[rune]int) bool {
	if len(labels) != 2 {
		return false
	}

	if containsValue(labels, 4) {
		return false
	}

	return true
}

func readHands(scanner *bufio.Scanner) hands {
	hands := []hand{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		cards := split[0]
		bid, _ := strconv.Atoi(split[1])
		hands = append(hands, hand{cards: cards, bid: bid, pokerHand: readPokerHand(cards)})
	}
	return hands
}

func processResult(hands hands) int {
	result := 0
	for i, hand := range hands {
		index := i + 1
		result += hand.bid * index
	}

	return result
}

func part2() (int, error) {
	cards = map[byte]Card{}
	cards['2'] = Two
	cards['3'] = Three
	cards['4'] = Four
	cards['5'] = Five
	cards['6'] = Six
	cards['7'] = Seven
	cards['8'] = Eight
	cards['9'] = Nine
	cards['T'] = Ten
	cards['J'] = Joker
	cards['Q'] = Queen
	cards['K'] = King
	cards['A'] = Ace

	file, err := os.Open("input1")
	if err != nil {
		fmt.Printf("Error reading file %v", file.Name())
	}

	defer file.Close()

	var result = 1
	sc := bufio.NewScanner(file)
	hands := readHands(sc)
	sort.Sort(hands)
	fmt.Println(hands)

	result = processResult(hands)
	return result, nil
}
