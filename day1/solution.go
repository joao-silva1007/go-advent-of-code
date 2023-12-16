package day1

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/joao-silva1007/go-advent-of-code/util"
)

func run() {
	strArray := util.ReadInput()
	fmt.Printf("day1 part1 result: %d\n", part1(strArray))
	fmt.Printf("day1 part2 result: %d\n", part2(strArray))
}

func part1(strArray []string) int {
	var first, last int
	sum := 0
	for _, line := range strArray {
		first = -1
		last = -1
		for _, r := range line {
			if unicode.IsDigit(r) {
				if first == -1 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}
		sum += first*10 + last
	}

	return sum
}

func part2(strArray []string) int {
	var first, last int
	sum := 0
	for _, line := range strArray {
		first = -1
		last = -1
		for i, r := range line {
			if unicode.IsDigit(r) {
				if first == -1 {
					first = int(r - '0')
				}
				last = int(r - '0')
			} else {
				found := getDigitsFromWord(line[i:])
				if found != -1 {
					if first == -1 {
						first = found
					}
					last = found
				}
			}
		}
		sum += first*10 + last
	}
	return sum
}

func getDigitsFromWord(entry string) int {
	wordToDigit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for word := range wordToDigit {
		if strings.Index(entry, word) == 0 {
			return wordToDigit[word]
		}
	}
	return -1
}
