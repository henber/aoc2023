package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadLines("day6/input.txt")
	parsedInput := parseInput(input)
	solveA(parsedInput)
	solveB(parsedInput)
}

type Input struct {
	times     []int
	distances []int
}

var numbersRegex = regexp.MustCompile("(\\d+)")

func parseInput(input []string) Input {
	return Input{
		times:     util.ConvertToInts(numbersRegex.FindAllString(input[0], -1)),
		distances: util.ConvertToInts(numbersRegex.FindAllString(input[1], -1)),
	}
}

func solveA(input Input) {
	var results []int
	for i := 0; i < len(input.times); i++ {
		winningOptions := playRace(input.times[i], input.distances[i])
		results = append(results, winningOptions)
	}
	prod := 1
	for _, result := range results {
		prod *= result
	}
	fmt.Printf("Answer A: %d\n", prod)
}

func playRace(time int, distance int) int {
	winningOptions := 0
	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			winningOptions += 1
		}
	}
	return winningOptions
}

func solveB(input Input) {
	time, _ := strconv.Atoi(strings.Join(util.ConvertToStrings(input.times), ""))
	distance, _ := strconv.Atoi(strings.Join(util.ConvertToStrings(input.distances), ""))

	fmt.Printf("Answer B: %d\n", playRace(time, distance))
}
