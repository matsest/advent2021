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

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(lines)
	//fmt.Println("Part 2: ", p2)
}
