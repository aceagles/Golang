package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.ReadFile("./input.txt")

	// map of all the coordinates of the symbols.
	symbols := map[image.Point]rune{}

	// Split on newlines and iterate through
	for y, row := range strings.Fields(string(input)) {
		for x, val := range row {
			if val != '.' && !unicode.IsDigit(val) {
				symbols[image.Point{x, y}] = val
			}
		}
	}

	part1, part2 := 0, 0
	adjacentNumbers := map[image.Point][]int{}
	re := regexp.MustCompile(`\d+`)
	for y, row := range strings.Fields(string(input)) {
		for _, ndxs := range re.FindAllStringIndex(row, -1) {
			boundaries := map[image.Point]struct{}{}
			for i := ndxs[0]; i < ndxs[1]; i++ {
				for _, v := range []image.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
					boundaries[image.Point{i, y}.Add(v)] = struct{}{}
				}
			}

			n, _ := strconv.Atoi(row[ndxs[0]:ndxs[1]])
			for val := range boundaries {
				_, ok := symbols[val]
				if ok {
					adjacentNumbers[val] = append(adjacentNumbers[val], n)
					part1 += n
				}
			}
		}

	}
	for loc, val := range symbols {
		if val == '*' {
			v := adjacentNumbers[loc]
			if len(v) == 2 {
				part2 += v[0] * v[1]
			}
		}
	}

	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)

}
