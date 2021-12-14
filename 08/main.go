package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/matsest/advent2021/utils"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func readNumber(s string) int {

	// cf len 2 -> 1
	// acf len 3 -> 7
	// bcdf len 4 -> 4
	// len 5 -> 2, 3, 5
	// len 6 -> 0, 6, 9
	// abcdefg len 7 -> 8

	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	return -1
}

func part1(entries []string) int {
	count := 0

	for _, e := range entries {
		digits := strings.Split(e, " ")
		for _, d := range digits {
			if readNumber(d) != 0 {
				count += 1
			}
		}
	}
	return count
}

type Entry struct {
	signalPattern []string
	outputValues  []string
}

func part2(entries []Entry) int {

	for _, e := range entries {
		for _, s := range e.signalPattern {
			if readNumber(s) != -1 {
				fmt.Println(s, "->", readNumber(s))
			} else if len(s) == 5 {
				fmt.Println(s, "->", "??? 2, 3 or 5")
			} else if (len(s)) == 6 {
				fmt.Println(s, "->", "??? 0, 6 or 9")
			}
		}
		fmt.Println(e)
	}

	return 0
}

func main() {
	lines, _ := utils.ReadLines("test.txt")

	// Parse input
	var outputvalueentries []string
	for _, v := range lines {
		parts := strings.Split(v, " | ")
		outputvalueentries = append(outputvalueentries, parts[1])
	}
	//fmt.Println(outputvalues)

	// Part 1
	p1 := part1(outputvalueentries)
	fmt.Println("Part 1: ", p1)

	// Parse input for part 2
	var entries []Entry
	for _, v := range lines {
		parts := strings.Split(v, " | ")
		signalpatterns := strings.Split(parts[0], " ")
		outputvalueentries := strings.Split(parts[1], " ")
		entries = append(entries, Entry{signalpatterns, outputvalueentries})
	}
	//fmt.Println(entries)

	// Part 2
	p2 := part2(entries)
	fmt.Println("Part 2: ", p2)
}
