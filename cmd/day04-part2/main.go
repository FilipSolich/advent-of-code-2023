package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	sum := 0
	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}
	for i, line := range lines {
		cardSplit := strings.Split(line, ": ")
		wins := []int{}
		gets := []int{}
		winPart := true
		for _, numStr := range strings.Fields(cardSplit[1]) {
			if numStr == "|" {
				winPart = false
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			if winPart {
				wins = append(wins, num)
			} else {
				gets = append(gets, num)
			}
		}

		winCount := 0
		for _, n := range gets {
			for _, w := range wins {
				if n == w {
					winCount++
				}
			}
		}

		for c := copies[i]; c > 0; c-- {
			for j := 1; j <= winCount; j++ {
				copies[i+j]++
			}
		}
	}

	for _, c := range copies {
		sum += c
	}

	fmt.Println(sum)
}
