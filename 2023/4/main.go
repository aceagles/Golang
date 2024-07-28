package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := float64(0)
	for scanner.Scan() {
		input := scanner.Text()

		afterColon := strings.Split(input, ":")[1]
		splitBar := strings.Split(afterColon, "|")
		winningNumbers, ourNumbers := strings.Fields(splitBar[0]), strings.Fields(splitBar[1])
		nWins := -1

		for _, val := range ourNumbers {
			if slices.Contains(winningNumbers, val) {
				nWins++
			}
		}
		if nWins >= 0 {
			points += math.Pow(2, float64(nWins))
		}
	}
	fmt.Println("Points: ", points)
}
