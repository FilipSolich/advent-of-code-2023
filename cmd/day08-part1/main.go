package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseLines(lines []string) map[string][2]string {
	m := map[string][2]string{}
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		dest := strings.Split(parts[1][1:][:len(parts[1])-2], ", ")
		m[parts[0]] = [2]string{dest[0], dest[1]}
	}
	return m
}

func main() {
	data, err := os.ReadFile("inputs/day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	input := lines[0]
	lines = lines[2:]
	nodes := parseLines(lines)
	start := "AAA"

	steps := 0
outer:
	for {
		for _, r := range input {
			s := string(r)
			if start == "ZZZ" {
				break outer
			}
			if s == "L" {
				start = nodes[start][0]
			} else {
				start = nodes[start][1]
			}
			steps++
		}
	}

	fmt.Println(steps)
}
