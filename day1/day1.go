package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const numbers = "0123456789"

func stripMap(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) >= 0 {
			return r
		}
		return -1
	}, str)
}

func solveA() {
	lines := util.ReadLines("day1/input.txt")

	sum := 0
	for _, str := range lines {
		numberString := stripMap(str, numbers)
		number, _ := strconv.Atoi(string(numberString[0]) + string(numberString[len(numberString)-1]))
		sum += number
	}
	fmt.Printf("Answer A: %d\n", sum)
}

type StringNumber struct {
	index  int
	number string
}

var numbersInString = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func solveB() {
	lines := util.ReadLines("day1/input.txt")

	sum := 0
	for _, str := range lines {
		var stringNumbers []StringNumber
		for i := 0; i < len(numbersInString); i++ {
			found := findStringNumbers(str, numbersInString[i], fmt.Sprintf("%d", i+1))
			stringNumbers = append(stringNumbers, found...)
		}
		foundNumeric := findNumbers(str)
		stringNumbers = append(stringNumbers, foundNumeric...)
		sort.SliceStable(stringNumbers, func(i, j int) bool {
			return stringNumbers[i].index < stringNumbers[j].index
		})
		number, _ := strconv.Atoi(stringNumbers[0].number + stringNumbers[len(stringNumbers)-1].number)
		sum += number
	}
	fmt.Printf("Answer B: %d\n", sum)
}

func findNumbers(str string) []StringNumber {
	re := regexp.MustCompile("[1-9]")
	matches := re.FindAllStringIndex(str, -1)
	var results []StringNumber
	for _, match := range matches {
		results = append(results, StringNumber{match[0], string(str[match[0]])})
	}
	return results
}

func findStringNumbers(str, findRegex, value string) []StringNumber {
	r := regexp.MustCompile(findRegex)
	matches := r.FindAllStringIndex(str, -1)
	var stringNumbers []StringNumber
	for _, match := range matches {
		stringNumbers = append(stringNumbers, StringNumber{match[0], value})
	}
	return stringNumbers
}

func main() {
	solveA()
	solveB()
}
