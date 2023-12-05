package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSpecial(char string) bool {
	return strings.Contains("*+&=@/%$#-", char)
}

func hasNeihborSymbol(lines []string, lineN int, i int) bool {
	if i > 0 {
		if isSpecial(string(lines[lineN][i-1])) {
			return true
		}
	}
	if i > 0 && lineN > 0 {
		if isSpecial(string(lines[lineN-1][i-1])) {
			return true
		}
	}
	if lineN > 0 {
		if isSpecial(string(lines[lineN-1][i])) {
			return true
		}
	}
	if i < len(lines[lineN])-1 && lineN > 0 {
		if isSpecial(string(lines[lineN-1][i+1])) {
			return true
		}
	}
	if i < len(lines[lineN])-1 {
		if isSpecial(string(lines[lineN][i+1])) {
			return true
		}
	}
	if i < len(lines[lineN])-1 && lineN < len(lines)-1 {
		if isSpecial(string(lines[lineN+1][i+1])) {
			return true
		}
	}
	if lineN < len(lines)-1 {
		if isSpecial(string(lines[lineN+1][i])) {
			return true
		}
	}
	if i > 0 && lineN < len(lines)-1 {
		if isSpecial(string(lines[lineN+1][i-1])) {
			return true
		}
	}

	return false
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
	for lineN, line := range lines {
		num := ""
		neighborSymbol := false
		for i, c := range line {
			char := string(c)
			if strings.Contains("0123456789", char) {
				num += char
				if !neighborSymbol {
					if hasNeihborSymbol(lines, lineN, i) {
						neighborSymbol = true
					}
				}
			}
			if i == len(line)-1 || !strings.Contains("0123456789", char) {
				if neighborSymbol {
					intNum, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					sum += intNum
				}
				num = ""
				neighborSymbol = false
			}
		}
	}

	fmt.Println(sum)
}
