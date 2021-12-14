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

func part2(lines []string, length int) int {

	// Initialize polymer
	seed := strings.Split(string(lines[0]), "")

	// Initialize pairs
	pairs := make(map[string]int)
	for i := 0; i < len(seed)-1; i++ {
		pairs[seed[i]+seed[i+1]] += 1
	}

	// Create map of pairInsertions
	pairInsertions := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		pairInsertions[parts[0]] = parts[1]
	}

	// Create a map of letter counts
	letterCount := make(map[string]int)
	for _, v := range seed {
		letterCount[v]+=1
	}

	// Grow polymers using map of pairs
	for j := 0; j < length; j++ {
		newPairs := make(map[string]int)
		for k, v := range pairs {
			new := pairInsertions[k]
			if v > 0 {
				newPairs[k] -= v
				newPairs[string(k[0])+new] += v
				newPairs[new+string(k[1])] += v
				letterCount[new] += v
			}
		}

		for k, v := range newPairs {
				pairs[k] += v
		}
	}

	// Create array of ints from counts
	numbers := make([]int, 0, len(letterCount))
	for _, v := range letterCount {
		numbers = append(numbers, v)
	}

	min, max := findMinAndMax(numbers)

	return max-min
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Part 1
	p1 := part1(lines, 10)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(lines, 40)
	fmt.Println("Part 2: ", p2)
}