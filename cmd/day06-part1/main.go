package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseInput(line string) []int {
	result := []int{}
	nums := strings.Fields(line)[1:]
	for _, num := range nums {
		result = append(result, strToInt(num))
	}
	return result
}

func main() {
	data, err := os.ReadFile("inputs/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := [][]int{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		input = append(input, parseInput(line))
	}

	result := 1
	for i := 0; i < len(input[0]); i++ {
		raceOpts := 0
		for j := 0; j <= input[0][i]; j++ {
			dist := j * (input[0][i] - j)
			if dist > input[1][i] {
				raceOpts++
			}
		}

		result *= raceOpts
	}

	fmt.Println(result)
}
