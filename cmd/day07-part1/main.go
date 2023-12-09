package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var kindValues = map[string]int{
	"five":      7,
	"four":      6,
	"fullhouse": 5,
	"three":     4,
	"twopair":   3,
	"onepair":   2,
	"highcard":  1,
}

type Hand struct {
	Cards string
	Bid   int
}

func (h Hand) Kind() int {
	var top int
	m := map[rune]int{}
	for _, card := range h.Cards {
		m[card]++
	}
	double := false
	triple := false
	for _, v := range m {
		switch v {
		case 5:
			top = kindValues["five"]
		case 4:
			top = max(top, kindValues["four"])
		case 3:
			if double {
				top = max(top, kindValues["fullhouse"])
			} else {
				top = max(top, kindValues["three"])
			}
			triple = true
		case 2:
			if double {
				top = max(top, kindValues["twopair"])
			} else if triple {
				top = max(top, kindValues["fullhouse"])
			} else {
				top = max(top, kindValues["onepair"])
			}
			double = true
		case 1:
			top = max(top, kindValues["highcard"])
		}
	}
	return top
}

func parseHand(line string) Hand {
	parts := strings.Fields(line)
	bid, _ := strconv.Atoi(parts[1])
	return Hand{Cards: parts[0], Bid: bid}
}

func sortFunc(a, b Hand) int {
	if a.Kind() != b.Kind() {
		return a.Kind() - b.Kind()
	}
	for i := 0; i < len(a.Cards); i++ {
		x, y := cardValues[string(a.Cards[i])], cardValues[string(b.Cards[i])]
		if x != y {
			return x - y
		}
	}
	return 0
}

func main() {
	data, err := os.ReadFile("inputs/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hands := []Hand{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		hands = append(hands, parseHand(line))
	}

	slices.SortFunc(hands, sortFunc)

	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.Bid
	}

	fmt.Println(result)
}
