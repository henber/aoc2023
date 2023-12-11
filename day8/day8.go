package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
)

func main() {
	input := util.ReadLines("day8/input.txt")
	instructions, lrMap := parseInput(input)

	solveA(instructions, lrMap)
	solveB(instructions, lrMap, input)
}

func solveA(instructions string, lrMap map[string]LeftRight) {
	steps := 0
	position := "AAA"

	for position != "ZZZ" {
		position = nextPosition(instructions, lrMap, steps, position)
		steps++
	}

	fmt.Printf("Answer A: %d", steps)
}

func solveB(instructions string, lrMap map[string]LeftRight, input []string) {
	steps := 0
	positions := findStartPositions(input[2:])

	var stepFactor []int

	for len(positions) > 0 {
		for i := 0; i < len(positions); i++ {
			positions[i] = nextPosition(instructions, lrMap, steps, positions[i])
		}
		steps++

		for i, position := range positions {
			if position[2] == 'Z' {
				stepFactor = append(stepFactor, steps)
				positions = append(positions[:i], positions[i+1:]...)
			}
		}
	}

	fmt.Printf("Answer B: %d", util.LCM(stepFactor[0], stepFactor[1], stepFactor[2:]...))
}

func findStartPositions(strings []string) []string {
	var startPos []string
	for _, s := range strings {
		if s[2] == 'A' {
			startPos = append(startPos, s[:3])
		}
	}
	return startPos
}

func nextPosition(instructions string, lrMap map[string]LeftRight, steps int, position string) string {
	instruction := instructions[steps%len(instructions)]
	if instruction == 'L' {
		return lrMap[position].left
	} else {
		return lrMap[position].right
	}
}

type LeftRight struct {
	left  string
	right string
}

var inputRegex = regexp.MustCompile("(\\w+)")

func parseInput(input []string) (string, map[string]LeftRight) {
	instructions := input[0]
	var leftRightMap = make(map[string]LeftRight, len(input))

	for i := 2; i < len(input); i++ {
		matches := inputRegex.FindAllString(input[i], -1)
		leftRightMap[matches[0]] = LeftRight{left: matches[1], right: matches[2]}
	}
	return instructions, leftRightMap
}
