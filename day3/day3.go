package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	input := util.ReadLines("day3/input.txt")
	solveA(input)
	solveB(input)
}

type Coord struct {
	x int
	y int
}

func (coord Coord) add(other Coord) Coord {
	return Coord{
		coord.x + other.x, coord.y + other.y,
	}
}

var scanOffsets = []Coord{
	{1, 0},
	//{0, 0},
	{-1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{1, 1},
	{0, 1},
	{-1, 1},
}

func scanSurroundingCoordinates(list []string, coord Coord, predicate func(uint8) bool) (bool, []Coord) {
	var foundCoords []Coord
	for _, offset := range scanOffsets {
		toCheck := coord.add(offset)
		if toCheck.y < 0 || toCheck.y >= len(list) {
			continue
		}
		if toCheck.x < 0 || toCheck.x >= len(list[toCheck.y]) {
			continue
		}
		char := list[toCheck.y][toCheck.x]
		if predicate(char) {
			foundCoords = append(foundCoords, toCheck)
		}
	}
	return len(foundCoords) > 0, foundCoords
}

func scanLeft(x int, str string, predicate func(uint8) bool) int {
	for i := x - 1; i >= 0; i-- {
		if !predicate(str[i]) {
			return i + 1
		}
	}
	return 0
}

func scanRight(x int, str string, predicate func(uint8) bool) int {
	for i := x + 1; i < len(str); i++ {
		if !predicate(str[i]) {
			return i
		}
	}
	return len(str)
}

func expandNumberString(x int, str string) string {
	isDigit := func(u uint8) bool {
		return unicode.IsDigit(rune(u))
	}
	return str[scanLeft(x, str, isDigit):scanRight(x, str, isDigit)]
}

func expandNumberIndexes(x int, str string) (int, int) {
	isDigit := func(u uint8) bool {
		return unicode.IsDigit(rune(u))
	}
	return scanLeft(x, str, isDigit), scanRight(x, str, isDigit)
}

func solveA(input []string) {
	var foundCoords []Coord
	for y, s := range input {
		prevWasHit := false
		for x, c := range s {
			if unicode.IsDigit(c) {
				coord := Coord{x, y}
				foundMatch, _ := scanSurroundingCoordinates(input, coord, func(u uint8) bool {
					char := rune(u)
					return u != '.' && !unicode.IsDigit(char) && !unicode.IsLetter(char)
				})
				if !prevWasHit && foundMatch {
					foundCoords = append(foundCoords, coord)
					prevWasHit = true
				}
			} else {
				prevWasHit = false
			}
		}
	}
	sum := 0
	for _, coord := range foundCoords {
		integer, _ := strconv.Atoi(expandNumberString(coord.x, input[coord.y]))
		sum += integer
	}
	fmt.Printf("Answer A: %d\n", sum)
}

type CoordRange struct {
	y  int
	x1 int
	x2 int
}

func containsRange(ranges []CoordRange, coord Coord) bool {
	for _, coordRange := range ranges {
		if coordRange.y == coord.y && coord.x >= coordRange.x1 && coord.x <= coordRange.x2 {
			return true
		}
	}
	return false
}

func coordsToRanges(coords []Coord, input []string) []CoordRange {
	var ranges []CoordRange
	for _, coord := range coords {
		if containsRange(ranges, coord) {
			continue
		}
		x1, x2 := expandNumberIndexes(coord.x, input[coord.y])
		ranges = append(ranges, CoordRange{coord.y, x1, x2})
	}
	return ranges
}

func (rang CoordRange) toNumber(input []string) int {
	integer, _ := strconv.Atoi(input[rang.y][rang.x1:rang.x2])
	return integer
}

func solveB(input []string) {
	sum := 0
	for y, s := range input {
		for x, c := range s {
			if c == '*' {
				coord := Coord{x, y}
				_, coords := scanSurroundingCoordinates(input, coord, func(u uint8) bool {
					return unicode.IsDigit(rune(u))
				})
				ranges := coordsToRanges(coords, input)
				if len(ranges) != 2 {
					continue
				}
				sum += ranges[0].toNumber(input) * ranges[1].toNumber(input)
			}
		}
	}

	fmt.Printf("Answer B: %d\n", sum)
}
