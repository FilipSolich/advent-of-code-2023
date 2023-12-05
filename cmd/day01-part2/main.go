package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNumStr(str string) string {
	for _, n := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		if str == n {
			return str
		}
	}
	switch str {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}

func main() {
	in, err := os.Open("inputs/day01/input.txt")
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
		var first, last string
		for i := 0; i < len(line); i++ {
			for j := i; j < len(line); j++ {
				num := getNumStr(line[i : j+1])
				if num != "" {
					if first == "" {
						first = num
					}
					last = num
				}
			}
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	fmt.Println(sum)
}
