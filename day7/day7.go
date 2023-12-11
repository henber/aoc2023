package main

import (
	"aoc2023/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadLines("day7/input.txt")
	solveA(input)
	solveB(input)
}

type Hand struct {
	cards    string
	bid      int
	handType int
}

const ( // iota is reset to 0
	HIGH_CARD  = iota
	ONE_PAIR   = iota
	TWO_PAIR   = iota
	THREE_KIND = iota
	FULL_HOUSE = iota
	FOUR_KIND  = iota
	FIVE_KIND  = iota
)

const (
	WILDCARD_JOKER = 0
	T              = 10
	J              = 11
	Q              = 12
	K              = 13
	A              = 14
)

func parseCard(rune rune, wildcardJoker bool) int {
	switch rune {
	case 'T':
		return T
	case 'J':
		if wildcardJoker {
			return WILDCARD_JOKER
		}
		return J
	case 'Q':
		return Q
	case 'K':
		return K
	case 'A':
		return A
	default:
		return int(rune - '0')
	}
}

func resolveHandType(hand Hand, jokerWildcard bool) int {
	cardMap := make(map[int]int, A)
	for _, c := range hand.cards {
		card := parseCard(c, jokerWildcard)
		cardMap[card] += 1
	}

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range cardMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i, kv := range ss {
		if kv.Key == WILDCARD_JOKER {
			continue
		}
		value := kv.Value
		if jokerWildcard {
			value += cardMap[WILDCARD_JOKER]
		}
		if value == 5 {
			return FIVE_KIND
		}
		if value == 4 {
			return FOUR_KIND
		}
		if value == 3 {
			if ss[i+1].Value == 2 {
				return FULL_HOUSE
			} else {
				return THREE_KIND
			}
		}
		if value == 2 {
			if ss[i+1].Value == 2 {
				return TWO_PAIR
			}
			return ONE_PAIR
		}
		return HIGH_CARD
	}
	return FIVE_KIND // "JJJJJ"
}

func parseInput(input []string, wildcardJoker bool) []Hand {
	var result []Hand
	for _, s := range input {
		split := strings.Split(s, " ")
		bid, _ := strconv.Atoi(split[1])
		hand := Hand{cards: split[0], bid: bid}
		hand.handType = resolveHandType(hand, wildcardJoker)
		result = append(result, hand)
	}

	return result
}

func rankHands(input []string, wildcardJoker bool) int {
	parsedInput := parseInput(input, wildcardJoker)
	sort.Slice(parsedInput, func(i, j int) bool {
		hand := parsedInput[i]
		other := parsedInput[j]
		if hand.handType != other.handType {
			return hand.handType > other.handType
		}
		for i := 0; i < len(hand.cards); i++ {
			if hand.cards[i] != other.cards[i] {
				return parseCard(rune(hand.cards[i]), wildcardJoker) > parseCard(rune(other.cards[i]), wildcardJoker)
			}
		}
		return false
	})

	sum := 0
	for i, hand := range parsedInput {
		sum += (len(input) - i) * hand.bid
	}
	return sum
}

func solveA(input []string) {
	sum := rankHands(input, false)
	fmt.Printf("Answer A: %d\n", sum)
}

func solveB(input []string) {
	sum := rankHands(input, true)
	fmt.Printf("Answer B: %d\n", sum)
}
