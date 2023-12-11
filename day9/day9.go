package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
)

func main() {
	input := parseInput(util.ReadLines("day9/input.txt"))

	solveA(input)
	solveB(input)
}

func solveA(input [][]int) {
	var results = make([]int, len(input))
	for i, ints := range input {
		results[i] = predictNext(ints)
	}

	fmt.Printf("Answer A: %d\n", util.Sum(results))
}

func solveB(input [][]int) {
	var results = make([]int, len(input))
	for i, ints := range input {
		results[i] = predictPrevious(ints)

	}
	fmt.Printf("Answer B: %d\n", util.Sum(results))
}

func predictNext(ints []int) int {
	if allEqual(ints) {
		return ints[0]
	}
	return ints[len(ints)-1] + predictNext(arrayDiff(ints))
}

func predictPrevious(ints []int) int {
	if allEqual(ints) {
		return ints[0]
	}
	return ints[0] - predictPrevious(arrayDiff(ints))
}

func arrayDiff(ints []int) []int {
	var results = make([]int, len(ints)-1)
	for i := 0; i < len(ints)-1; i++ {
		results[i] = ints[i+1] - ints[i]
	}
	return results
}

func allEqual(ints []int) bool {
	first := ints[0]
	for i := 1; i < len(ints); i++ {
		if first != ints[i] {
			return false
		}
	}
	return true
}

var inputRegex = regexp.MustCompile("(-*\\d+)")

func parseInput(input []string) [][]int {
	var parsedInput = make([][]int, len(input))

	for i, s := range input {
		result := inputRegex.FindAllString(s, -1)
		parsedInput[i] = util.ConvertToInts(result)
	}

	return parsedInput
}
