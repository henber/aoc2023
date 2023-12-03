package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadLines("day2/input.txt")
	gameRounds := parseInput(input)
	solveA(gameRounds)
	solveB(gameRounds)
}

type BallRound struct {
	red   int
	green int
	blue  int
}

func parseInput(input []string) [][]BallRound {
	var gameRounds [][]BallRound
	for _, s := range input {
		gameRoundIndex := strings.Index(s, ":")
		s = s[gameRoundIndex+1:]
		rounds := strings.Split(s, ";")
		gameRounds = append(gameRounds, parseRounds(rounds))
	}
	return gameRounds
}

var ballRegex = regexp.MustCompile("([0-9]+ \\w+)")

func parseRounds(rounds []string) []BallRound {
	var resultBalls []BallRound
	for _, round := range rounds {
		balls := ballRegex.FindAllString(round, -1)
		ballRound := BallRound{}
		for _, ball := range balls {
			components := strings.Split(ball, " ")
			nr, _ := strconv.Atoi(components[0])
			switch components[1] {
			case "blue":
				ballRound.blue = nr
			case "red":
				ballRound.red = nr
			case "green":
				ballRound.green = nr
			}
		}
		resultBalls = append(resultBalls, ballRound)
	}
	return resultBalls
}

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func isRoundOk(round []BallRound) bool {
	for _, ballRound := range round {
		if ballRound.green > maxGreen || ballRound.blue > maxBlue || ballRound.red > maxRed {
			return false
		}
	}
	return true
}

func solveA(input [][]BallRound) {
	sum := 0

	for i, round := range input {
		if isRoundOk(round) {
			sum += i + 1
		}
	}

	fmt.Printf("Answer A: %d\n", sum)
}

func collapseToMin(rounds []BallRound) BallRound {
	min := rounds[0]
	for i := 1; i < len(rounds); i++ {
		round := rounds[i]
		if round.green > min.green {
			min.green = round.green
		}
		if round.red > min.red {
			min.red = round.red
		}
		if round.blue > min.blue {
			min.blue = round.blue
		}
	}
	return min
}

func ballRoundProduct(round BallRound) int {
	return round.green * round.blue * round.red
}

func solveB(input [][]BallRound) {
	sum := 0
	for _, rounds := range input {
		minRound := collapseToMin(rounds)
		sum += ballRoundProduct(minRound)
	}
	fmt.Printf("Answer B: %d", sum)
}
