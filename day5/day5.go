package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

func main() {
	input := util.ReadLines("day5/input.txt")
	seeds, mappers := parseInput(input)
	solveA(seeds, mappers)
	start := time.Now()
	solveB(seeds, mappers)
	elapsed := time.Since(start)
	fmt.Printf("Execution for B took %s", elapsed)
}

func minFind(arr []int) int {
	min := arr[0]
	for _, num1 := range arr {
		if num1 < min {
			min = num1
		}
	}
	return min
}

func solveA(seeds []int, mapperChain [][]Mapper) {
	var results []int
	for _, seed := range seeds {
		results = append(results, performMapChain(seed, mapperChain))
	}
	fmt.Printf("Answer A: %d\n", minFind(results))
}

func solveB(seeds []int, mappers [][]Mapper) {
	ranges := parseRanges(seeds)
	fmt.Println(len(ranges))

	result := performRangeMapChain(ranges, mappers)
	fmt.Println(len(result))

	min := math.MaxInt64
	for _, res := range result {
		//fmt.Println(res)
		if res.start < min {
			min = res.start
		}
	}

	fmt.Printf("Answer B: %d\n", min)
}

type Range struct {
	start int
	end   int
}

func parseRanges(seeds []int) []Range {
	var result []Range
	for i := 0; i < len(seeds); i += 2 {
		result = append(result, Range{seeds[i], seeds[i] + seeds[i+1]})
	}
	return result
}

func performMapChain(seed int, mapChain [][]Mapper) int {
	updatedSeed := seed
	for _, mappers := range mapChain {
		updatedSeed = mapSeed(updatedSeed, mappers)
	}
	return updatedSeed
}

func performRangeMapChain(ranges []Range, mapChain [][]Mapper) []Range {
	updatedRanges := ranges
	for _, mappers := range mapChain {
		updatedRanges = mapRanges(updatedRanges, mappers)
	}
	return updatedRanges
}

func applyMapper(nr int, mapper Mapper) int {
	return mapper.r1 + (nr - mapper.r2)
}

func mapRanges(ranges []Range, mappers []Mapper) []Range {
	var mappedRanges []Range
	for _, r := range ranges {
		mappedRanges = append(mappedRanges, mapRange(r, mappers)...)
	}
	return mappedRanges
}

func mapRange(r Range, mappers []Mapper) []Range {
	var unmappedRanges []Range
	var mappedRanges []Range
	for _, mapper := range mappers {
		// included fully
		if r.start >= mapper.r2 && r.end <= mapper.r2+mapper.size {
			mappedRanges = append(mappedRanges, Range{applyMapper(r.start, mapper), applyMapper(r.end, mapper)})
			break
		}
		// partial overlap from left
		if r.start < mapper.r2 && r.end >= mapper.r2 && r.end <= mapper.r2+mapper.size {
			lower := Range{r.start, mapper.r2 - 1}
			upper := Range{mapper.r2, r.end}
			unmappedRanges = append(unmappedRanges, lower, upper)
			break
		}
		// partial overlap from right
		if r.end > mapper.r2+mapper.size && r.start >= mapper.r2 && r.start <= mapper.r2+mapper.size {
			upper := Range{mapper.r2 + mapper.size + 1, r.end}
			lower := Range{r.start, mapper.r2 + mapper.size}
			unmappedRanges = append(unmappedRanges, lower, upper)
			break
		}
		// fully overlapping
		if r.start < mapper.r2 && r.end > mapper.r2+mapper.size {
			lower := Range{r.start, mapper.r2 - 1}
			upper := Range{mapper.r2 + mapper.size + 1, r.end}
			unmappedRanges = append(unmappedRanges, lower, upper)
			mappedRanges = append(mappedRanges, Range{applyMapper(mapper.r2, mapper), applyMapper(mapper.r2+mapper.size, mapper)})
			break
		} else {
			// no overlap
		}
	}

	if len(unmappedRanges) == 0 && len(mappedRanges) == 0 {
		return []Range{r}
	}
	// recursively map the remaining unmapped ranges
	mappedRanges = append(mappedRanges, mapRanges(unmappedRanges, mappers)...)
	return mappedRanges
}

func mapSeed(seed int, mappers []Mapper) int {
	for _, mapper := range mappers {
		if seed >= mapper.r2 && seed <= mapper.r2+mapper.size {
			return mapper.r1 + (seed - mapper.r2)
		}
	}
	return seed
}

var numbersRegex = regexp.MustCompile("(\\d+)")

func parseInput(input []string) ([]int, [][]Mapper) {
	seeds := util.ConvertToInts(numbersRegex.FindAllString(input[0], -1))

	var mappers [][]Mapper
	for i := 2; i < len(input); i++ {
		if len(numbersRegex.FindAllString(input[i], -1)) == 0 && len(strings.TrimSpace(input[i])) > 0 {
			newIndex, parsedMappers := parseMappers(i+1, input)
			mappers = append(mappers, parsedMappers)
			i = newIndex
		} else if len(strings.TrimSpace(input[i])) == 0 {
			continue
		}
	}
	return seeds, mappers
}

func parseMappers(index int, input []string) (int, []Mapper) {
	var mappers []Mapper
	for i := index; i < len(input); i++ {
		if len(strings.TrimSpace(input[i])) == 0 {
			return i, mappers
		}
		numbers := util.ConvertToInts(numbersRegex.FindAllString(input[i], -1))
		mappers = append(mappers, Mapper{numbers[0], numbers[1], numbers[2]})
	}
	return len(input), mappers
}

type Mapper struct {
	r1   int
	r2   int
	size int
}
