package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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
		first, last := "", ""
		for _, char := range line {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			if first == "" {
				first = string(char)
			}
			last = string(char)
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	fmt.Println(sum)
}
