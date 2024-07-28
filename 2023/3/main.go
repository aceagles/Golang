package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func toNumbersSlice(s []string) []int {
	var numbers []int
	for _, val := range s {
		num, err := strconv.ParseInt(val, 10, 32)
		if err == nil {
			numbers = append(numbers, int(num))
		}
	}
	return numbers
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func searchAlongString(indexes []int, s string, l int) bool {
	prevRune := []rune(s)
	if len(prevRune) == 0 {
		return false
	}
	for i := indexes[0] - 1; i <= indexes[1]; i++ {
		if i >= 0 && i < l {
			if isSymbol(prevRune[i]) {
				return true
			}
		}
	}
	return false

}
func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prevLine, currLine, nextLine string
	re := regexp.MustCompile(`\d+`)

	runAgain := true
	numberSlice := []int{}
	for runAgain {
		runAgain = scanner.Scan()
		nextLine = scanner.Text()
		if currLine != "" {
			lineLength := len(currLine)
			numberStrings := re.FindAllString(currLine, -1)
			numbers := toNumbersSlice(numberStrings)
			indexs := re.FindAllStringIndex(currLine, -1)
			for j, ndx := range indexs {
				if searchAlongString(ndx, prevLine, lineLength) || searchAlongString(ndx, currLine, lineLength) || searchAlongString(ndx, nextLine, lineLength) {
					numberSlice = append(numberSlice, numbers[j])
				}
			}
		}

		prevLine, currLine = currLine, nextLine
	}
	fmt.Println(numberSlice)
	sum := 0
	for _, val := range numberSlice {
		sum += val
	}
	fmt.Print("Sum: ", sum)
}
