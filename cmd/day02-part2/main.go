package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

		sets := strings.Split(gameLine[1], "; ")
		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				var num int
				var colorStr string
				fmt.Sscanf(color, "%d %s", &num, &colorStr)
				switch colorStr {
				case "red":
					if num > maxRed {
						maxRed = num
					}
				case "green":
					if num > maxGreen {
						maxGreen = num
					}
				case "blue":
					if num > maxBlue {
						maxBlue = num
					}
				}
			}
		}
		gamePower := maxRed * maxGreen * maxBlue
		sum += gamePower
	}

	fmt.Println(sum)
}
