package util

import (
	"bufio"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ConvertToInts(array []string) []int {
	var res []int
	for _, s := range array {
		integer, _ := strconv.Atoi(s)
		res = append(res, integer)
	}
	return res
}

func ConvertToStrings(array []int) []string {
	var res []string
	for _, s := range array {
		s := strconv.Itoa(s)
		res = append(res, s)
	}
	return res
}
