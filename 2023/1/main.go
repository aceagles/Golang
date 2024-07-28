package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numberStrings = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func convertNumericStrings(s string) string {
	for key, value := range numberStrings {
		fmt.Println(key)
		s = strings.Replace(s, key, value, -1)
	}
	return s
}

func reverseRunes(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int64
	i := 0
	for scanner.Scan() {
		i++
		text := scanner.Text()
		fmt.Println(text)
		convertedText := convertNumericStrings(text)
		// fmt.Println(convertedText)
		runes := []rune(convertedText)
		var a, b rune

		for _, r := range runes {
			if unicode.IsDigit(r) {
				a = r
				break
			}

		}
		reverseRunes(runes)
		for _, r := range runes {
			if unicode.IsDigit(r) {
				b = r
				break
			}
		}
		calbrationValue := string(a) + string(b)

		x, _ := strconv.ParseInt(calbrationValue, 10, 32)
		fmt.Println(x)
		sum += x
	}
	fmt.Println(i)
	fmt.Println(sum)
}
