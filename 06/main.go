package main

import (
	"fmt"

	"github.com/matsest/advent2021/utils"
)

func part1(seeds []int) int {

	lanternfishes := make([]int, len(seeds))
	copy(lanternfishes, seeds)

	for days := 0; days < 80; days++ {
		var tmp []int
		for i := 0; i < len(lanternfishes); i++ {
			if lanternfishes[i] == 0 {
				lanternfishes[i] = 6
				tmp = append(tmp, 8)
			} else {
				lanternfishes[i]--
			}
		}
		lanternfishes = append(lanternfishes, tmp...)
	}

	return len(lanternfishes)
}

func part2(seeds []int) int {

	lanternfishes := make([]int, len(seeds))
	copy(lanternfishes, seeds)

	// reduce to map to remove iteration cost
	var lanternfishCount = make(map[int]int)

	for i := range lanternfishes {
		lanternfishCount[lanternfishes[i]]++
	}

	for day := 0; day < 256; day++ {
		var newCount = map[int]int{}
		for k, v := range lanternfishCount {
			if k == 0 {
				newCount[6] += v
				newCount[8] += v
			} else {
				newCount[k-1] += v
			}
		}
		lanternfishCount = newCount
	}

	sum := 0
	for _, v := range lanternfishCount {
		sum += v
	}
	return sum
}

func main() {
	lines, _ := utils.ReadInts("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(lines)
	fmt.Println("Part 2: ", p2)
}
