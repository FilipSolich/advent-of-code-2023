package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

func main() {
	in, err := os.Open("inputs/day02/input.txt")
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
	for _, line := range lines {
		gameLine := strings.Split(line, ": ")
		var gameID int
		fmt.Sscanf(gameLine[0], "Game %d", &gameID)

		gameFitsIn := true
		sets := strings.Split(gameLine[1], "; ")
	outer:
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				var num int
				var colorStr string
				fmt.Sscanf(color, "%d %s", &num, &colorStr)
				if (colorStr == "red" && num > MaxRed) || (colorStr == "green" && num > MaxGreen) || (colorStr == "blue" && num > MaxBlue) {
					gameFitsIn = false
					break outer
				}
			}
		}
		if gameFitsIn {
			sum += gameID
		}
	}

	fmt.Println(sum)
}
