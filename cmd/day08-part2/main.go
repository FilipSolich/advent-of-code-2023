package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func walkThrough(input string, nodes map[string][2]string, start string, out chan<- int) {
	steps := 0
	for {
		for _, r := range input {
			s := string(r)
			if strings.HasSuffix(start, "Z") {
				out <- steps
			}
			if s == "L" {
				start = nodes[start][0]
			} else {
				start = nodes[start][1]
			}
			steps++
		}
	}
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

	starts := []string{}
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	chs := []chan int{}
	fetch := []bool{}
	for range starts {
		chs = append(chs, make(chan int, 1))
		fetch = append(fetch, true)
	}
	for n, s := range starts {
		go walkThrough(input, nodes, s, chs[n])
	}

	steps := 0
	res := make([]int, len(starts))
outer:
	for {
		for n, v := range fetch {
			if v {
				res[n] = <-chs[n]
			}
			fetch[n] = true
		}
		fmt.Println(res)
		if len(slices.Compact(res)) == 1 {
			steps = res[0]
			break outer
		}
		for n := range fetch {
			if res[n] == slices.Max(res) {
				fetch[n] = false
			}
		}
	}

	//	steps := 0
	//outer:
	//	for {
	//		for _, r := range input {
	//			s := string(r)
	//			end := 0
	//			for i, n := range starts {
	//				if strings.HasSuffix(n, "Z") {
	//					end++
	//				}
	//				if s == "L" {
	//					starts[i] = nodes[n][0]
	//				} else {
	//					starts[i] = nodes[n][1]
	//				}
	//			}
	//			if end == len(starts) {
	//				break outer
	//			}
	//			steps++
	//		}
	//	}

	fmt.Println(steps)
}
