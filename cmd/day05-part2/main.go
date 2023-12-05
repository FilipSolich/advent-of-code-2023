package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Source      int
	Destination int
	Range       int
}

type Map struct {
	From   string
	To     string
	Ranges []Range
}

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseSeeds(data string) [][]int {
	seeds := [][]int{}
	seedBlock := strings.Fields(strings.TrimPrefix(data, "seeds: "))
	for i := 0; i < len(seedBlock)/2; i++ {
		src := strToInt(seedBlock[i*2])
		rang := strToInt(seedBlock[i*2+1])
		seeds = append(seeds, []int{src, rang})
	}

	return seeds
}

func parseMap(data string) Map {
	m := Map{Ranges: []Range{}}

	lines := strings.Split(data, "\n")
	mapName := strings.Split(strings.TrimSuffix(lines[0], " map:"), "-to-")
	m.From = mapName[0]
	m.To = mapName[1]

	for _, line := range lines[1:] {
		nums := strings.Fields(line)
		r := Range{}
		r.Destination = strToInt(nums[0])
		r.Source = strToInt(nums[1])
		r.Range = strToInt(nums[2])
		m.Ranges = append(m.Ranges, r)
	}

	return m
}

func findDest(src int, m Map) int {
	for _, r := range m.Ranges {
		if src >= r.Source && src < r.Source+r.Range {
			delta := src - r.Source
			return r.Destination + delta
		}
	}
	return src
}

func processRange(r []int, maps []Map, resultCh chan<- int) {
	min := math.MaxInt
	for i := 0; i < r[1]; i++ {
		x := r[0] + i
		for _, m := range maps {
			x = findDest(x, m)
		}
		if x < min {
			min = x
		}
	}
	resultCh <- min
}

func main() {
	data, err := os.ReadFile("inputs/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	blocks := strings.Split(string(data), "\n\n")

	seeds := parseSeeds(blocks[0])
	maps := []Map{}
	for _, block := range blocks[1:] {
		m := parseMap(block)
		maps = append(maps, m)
	}

	mins := make(chan int, len(seeds))
	for _, seedRange := range seeds {
		go processRange(seedRange, maps, mins)
	}

	minSlice := []int{}
	for i := 0; i < len(seeds); i++ {
		minSlice = append(minSlice, <-mins)
	}

	fmt.Println(slices.Min(minSlice))
}
