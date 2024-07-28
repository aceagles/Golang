package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type valMap struct {
	destination, source, rng int
}

func mapVal(val int, mp valMap) int {
	if val >= mp.source && val <= mp.source+mp.rng-1 {
		fmt.Printf("Value %d Source %d Dest %d = ", val, mp.source, mp.destination)
		return val - mp.source + mp.destination
	}
	return val
}

func main() {
	input, _ := os.ReadFile("./input.txt")

	lines := strings.Split(string(input), "\n")
	var seeds []int
	for _, val := range strings.Fields(lines[0]) {
		intVal, err := strconv.Atoi(val)
		if err == nil {
			seeds = append(seeds, intVal)
		}
	}

	r := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	var mappings [][]valMap
	for _, val := range lines[1:] {
		if val != "" {
			matches := r.FindAllSubmatch([]byte(val), -1)
			if matches != nil {
				mp := valMap{}
				intVal, _ := strconv.Atoi(string(matches[0][1]))
				mp.destination = intVal
				intVal, _ = strconv.Atoi(string(matches[0][2]))
				mp.source = intVal
				intVal, _ = strconv.Atoi(string(matches[0][3]))
				mp.rng = intVal
				mappings[len(mappings)-1] = append(mappings[len(mappings)-1], mp)
			} else {
				mappings = append(mappings, []valMap{})
			}
		}
	}

	locVals := []int{}
	for _, seedVal := range seeds {
		fmt.Println("init ", seedVal)
		for _, mpGroup := range mappings {
			for _, mpVal := range mpGroup {
				prevVal := seedVal
				seedVal = mapVal(seedVal, mpVal)
				if prevVal != seedVal {
					fmt.Println(seedVal)
					break
				}
			}
		}
		locVals = append(locVals, seedVal)
		fmt.Println("-----")
	}

	fmt.Println(locVals)
	fmt.Println(slices.Min(locVals))
}
