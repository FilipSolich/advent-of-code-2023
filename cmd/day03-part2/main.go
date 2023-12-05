package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func isGear(char string) bool {
	return strings.Contains("*", char)
}

func hasNeihborSymbol(lines []string, lineN int, i int) (bool, int, int) {
	if i > 0 {
		if isGear(string(lines[lineN][i-1])) {
			return true, lineN, i - 1
		}
	}
	if i > 0 && lineN > 0 {
		if isGear(string(lines[lineN-1][i-1])) {
			return true, lineN - 1, i - 1
		}
	}
	if lineN > 0 {
		if isGear(string(lines[lineN-1][i])) {
			return true, lineN - 1, i
		}
	}
	if i < len(lines[lineN])-1 && lineN > 0 {
		if isGear(string(lines[lineN-1][i+1])) {
			return true, lineN - 1, i + 1
		}
	}
	if i < len(lines[lineN])-1 {
		if isGear(string(lines[lineN][i+1])) {
			return true, lineN, i + 1
		}
	}
	if i < len(lines[lineN])-1 && lineN < len(lines)-1 {
		if isGear(string(lines[lineN+1][i+1])) {
			return true, lineN + 1, i + 1
		}
	}
	if lineN < len(lines)-1 {
		if isGear(string(lines[lineN+1][i])) {
			return true, lineN + 1, i
		}
	}
	if i > 0 && lineN < len(lines)-1 {
		if isGear(string(lines[lineN+1][i-1])) {
			return true, lineN + 1, i - 1
		}
	}

	return false, 0, 0
}

func main() {
	in, err := os.Open("inputs/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	data, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	datas := string(data)
	lines := strings.Split(datas, "\n")

	sum := 0
	gearsMap := map[string][]int{}
	for lineN, line := range lines {
		num := ""
		key := ""
		neighborGear := false
		for i, c := range line {
			char := string(c)
			if strings.Contains("0123456789", char) {
				num += char
				if !neighborGear {
					neighbor, x, y := hasNeihborSymbol(lines, lineN, i)
					if neighbor {
						neighborGear = true
						key = strconv.Itoa(x) + ":" + strconv.Itoa(y)
					}
				}
			}
			if i == len(line)-1 || !strings.Contains("0123456789", char) {
				if neighborGear {
					intNum, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}

					if gearsMap[key] == nil {
						gearsMap[key] = []int{}
					}
					gearsMap[key] = append(gearsMap[key], intNum)
				}
				num = ""
				neighborGear = false
			}
		}
	}

	for _, v := range gearsMap {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}

	fmt.Println(sum)
}
