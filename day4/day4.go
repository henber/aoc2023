package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadLines("day4/input.txt")
	scoreCards := parseInput(input)
	solveA(scoreCards)
	solveB(scoreCards)
}

func solveA(cards []ScoreCard) {
	sum := 0
	for _, card := range cards {
		winners := score(card)
		if len(winners) <= 0 {
			continue
		}
		sum += int(math.Pow(float64(2), float64(len(winners)-1)))
	}
	fmt.Printf("Answer A: %d\n", sum)
}

func solveB(cards []ScoreCard) {
	result := make(map[int]int, len(cards))
	for i := 1; i <= len(cards); i++ {
		result[i] = 1
	}

	for i, card := range cards {
		winners := score(card)
		if len(winners) <= 0 {
			continue
		}
		for j := 1; j <= len(winners); j++ {
			result[i+1+j] += result[i+1]
		}
	}
	sum := 0
	for _, value := range result {
		sum += value
	}
	fmt.Printf("Answer B: %d\n", sum)
}

func score(card ScoreCard) []int {
	var winningNumbers []int
	for _, number := range card.numbers {
		_, ok := card.winners[number]
		if ok {
			winningNumbers = append(winningNumbers, number)
		}
	}
	return winningNumbers
}

type ScoreCard struct {
	winners map[int]bool
	numbers []int
}

var numberRegex = regexp.MustCompile("(\\d+)")

func convertToIntList(list []string) []int {
	var result []int
	for _, s := range list {
		integer, _ := strconv.Atoi(s)
		result = append(result, integer)
	}
	return result
}

func convertToIntSet(list []int) map[int]bool {
	result := map[int]bool{}
	for _, value := range list {
		result[value] = true
	}
	return result
}

func parseInput(input []string) []ScoreCard {
	var result []ScoreCard
	for _, s := range input {
		gameRoundIndex := strings.Index(s, ":")
		s = s[gameRoundIndex+1:]
		winnerSeparatorIndex := strings.Index(s, "|")
		winners := numberRegex.FindAllString(s[:winnerSeparatorIndex], -1)
		numbers := numberRegex.FindAllString(s[winnerSeparatorIndex:], -1)
		result = append(result, ScoreCard{convertToIntSet(convertToIntList(winners)), convertToIntList(numbers)})
	}
	return result
}
