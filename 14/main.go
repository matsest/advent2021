package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2021/utils"
)

func part1(lines []string, length int) int {

	seed := strings.Split(string(lines[0]), "")

	// Create map of pairInsertions
	pairInsertions := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		from := parts[0]
		to := parts[2]
		pairInsertions[from] = to
	}

	// Grow polymers
	var ans []string
	for j := 0; j < length; j++ {

		var tmp []string
		for i := 0; i < len(seed)-1; i++ {
			in := seed[i] + seed[i+1]
			out := pairInsertions[in]
			if i == 0 {
				tmp = append(tmp, seed[i])
			}
			tmp = append(tmp, seed[i+1:i+1]...)
			tmp = append(tmp, out)
			tmp = append(tmp, seed[i+1:i+2]...)
		}

		seed = tmp
		ans = tmp
	}

	// create count dictionary
	counts := make(map[string]int)

	for _, v := range ans {
		counts[v] += 1
	}

	// create array of ints from counts
	numbers := make([]int, 0, len(counts))
	for _, v := range counts {
		numbers = append(numbers, v)
	}

	min, max := findMinAndMax(numbers)

	return max - min
}

// Helper function to find max and min of an array
func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	lines, _ := utils.ReadLines("test.txt")

	// Part 1
	p1 := part1(lines, 10)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part1(lines, 10)
	fmt.Println("Part 2: ", p2)
}
