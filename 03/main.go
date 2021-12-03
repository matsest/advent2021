package main

import (
	"fmt"
	"strconv"

	"github.com/matsest/advent2021/utils"
)

func GammaEntry(m map[int]int) (ans string) {
	keys := make([]int, 0, len(m))
	values := make([]int, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	if values[0] > values[1] {
		ans = strconv.Itoa(keys[0])
	} else {
		ans = strconv.Itoa(keys[1])
	}
	return ans
}

func EpsilonEntry(m map[int]int) (ans string) {
	keys := make([]int, 0, len(m))
	values := make([]int, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	if values[0] > values[1] {
		ans = strconv.Itoa(keys[1])
	} else {
		ans = strconv.Itoa(keys[0])
	}
	return ans
}

func part1(diagnostics []string) int64 {
	var gamma string
	var epsilon string

	// Go through each column
	for i := 0; i < len(diagnostics[0]); i++ {
		var values []string
		for _, line := range diagnostics {
			values = append(values, string(line[i]))
		}
		counter := map[int]int{0: 0, 1: 0}

		// Go through each line in current column
		for _, val := range values {
			intVal, _ := strconv.Atoi(val)
			counter[intVal] += 1
		}

		// Find the current entry
		gamma += GammaEntry(counter)
		epsilon += EpsilonEntry(counter)
	}

	// Convert from Binary to Int
	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaDec * epsilonDec
}

func rating(diagnostics []string, ratingType string, i int) int64 {

	// Go through each line in i'th column
	counter := map[int]int{0: 0, 1: 0}
	var ones []string
	var zeros []string
	for _, line := range diagnostics {
		current := string(line[i])
		intVal, _ := strconv.Atoi(current)
		counter[intVal] += 1
		if current == "1" {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}

	var newDiag []string
	// oxy
	if ratingType == "oxy" {
		if len(zeros) > len(ones) {
			newDiag = zeros
		} else {
			newDiag = ones
		}
	} else {
		// co2
		if len(zeros) > len(ones) {
			newDiag = ones
		} else {
			newDiag = zeros
		}
	}

	if len(newDiag) == 1 {
		// Convert from Binary to Int
		ans, _ := strconv.ParseInt(string(newDiag[0]), 2, 64)
		return ans
	}

	return rating(newDiag, ratingType, i+1)
}

func part2(diagnostics []string) int64 {
	return rating(diagnostics, "oxy", 0) * rating(diagnostics, "co2", 0)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(lines)
	fmt.Println("Part 2: ", p2)
}