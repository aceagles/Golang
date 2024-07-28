package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var testMap = map[string]int{
	"blue":  14,
	"red":   12,
	"green": 13,
}

func parseGameNumber(s string) int64 {
	r := regexp.MustCompile(`Game (\d+).*`)

	matches := r.FindStringSubmatch(s)
	match, _ := strconv.ParseInt(matches[1], 10, 32)
	return match
}

func getCubes(s string) (int, string) {
	r := regexp.MustCompile(`(\d+) ([a-z]+)`)
	matches := r.FindStringSubmatch(s)
	intMatch, _ := strconv.ParseInt(matches[1], 10, 32)
	return int(intMatch), matches[2]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		texts := strings.Split(text, ":")
		guesses := strings.Split(texts[1], ", ")
		var individualPicks []string

		maxColour := map[string]int{}
		for _, val := range guesses {
			tmp := strings.Split(val, "; ")
			individualPicks = append(individualPicks, tmp...)
		}
		for _, val := range individualPicks {
			num, colour := getCubes(val)
			if num > maxColour[colour] {
				maxColour[colour] = num
			}
		}
		power := 1
		for _, val := range maxColour {
			power *= val
		}
		sum += power
	}
	fmt.Println(sum)
}
