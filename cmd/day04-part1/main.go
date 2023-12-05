package main

import (
	"fmt"
	"log"
	"math"
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
	for _, line := range lines {
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

		if winCount > 0 {
			sum += int(math.Pow(2, float64(winCount-1)))
		}
	}

	fmt.Println(sum)
}
