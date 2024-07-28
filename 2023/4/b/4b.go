package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input.txt")

	lines := strings.Split(string(input), "\n")

	cards := map[int]int{}
	for i, input := range lines {
		cards[i]++
		afterColon := strings.Split(input, ":")[1]
		splitBar := strings.Split(afterColon, "|")
		winningNumbers, ourNumbers := strings.Fields(splitBar[0]), strings.Fields(splitBar[1])
		nWins := 0

		for _, val := range ourNumbers {
			if slices.Contains(winningNumbers, val) {
				nWins++
			}
		}
		for n := 0; n < nWins; n++ {
			cards[i+1+n] += cards[i]
		}
	}

	sum := 0
	for _, val := range cards {
		sum += val
	}
	fmt.Println("Sum: ", sum)

}
