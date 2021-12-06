package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2021/utils"
)

type BingoEntry struct {
	number int
	marked bool
}

type Board struct {
	numbers [][]BingoEntry
	bingo   bool
}

func parseBingo(bingo []string) (numbers []int, boards []Board) {

	numbersString := bingo[0]
	numbersStringArray := strings.Split(numbersString, ",")
	//fmt.Println("Numbers: ", numbersStringArray)

	numbers, _ = utils.SliceAtoi(numbersStringArray)

	var b = Board{[][]BingoEntry{}, false}
	for _, row := range bingo[2:] {
		bingoRowString := strings.Fields(row)
		//fmt.Println("Bingo row string", bingoRowString)
		bingoRow, _ := utils.SliceAtoi(bingoRowString)
		bingoEntryRow := []BingoEntry{}
		for _, v := range bingoRow {
			bingoEntryRow = append(bingoEntryRow, BingoEntry{v, false})
		}

		//fmt.Println("Bingo row", bingoRow)
		// Separate boards by empty line
		if len(row) == 0 {
			//fmt.Println("NYTT")
			boards = append(boards, b)
			b = Board{[][]BingoEntry{}, false}
		} else {
			//fmt.Println(row)
			b.numbers = append(b.numbers, bingoEntryRow)
		}

		// append last board
		if row == bingo[len(bingo)-1] {
			boards = append(boards, b)
		}
	}

	//for i, bs := range boards {
	//	fmt.Println("Board ", i, bs)
	//}
	return numbers, boards
}

func checkBingo(board Board) (win bool, winningLine []BingoEntry) {
	// check rows
	for _, row := range board.numbers {
		bingoLength := len(row)
		markings := 0
		for _, entry := range row {
			if entry.marked {
				markings += 1
			}
			if markings == bingoLength {
				//fmt.Println("ROW WINNER")
				return true, row
			}
		}
	}
	// check columns
	for i, _ := range board.numbers[0] {
		bingoLength := len(board.numbers[0])
		markings := 0
		resCol := []BingoEntry{}
		for j, _ := range board.numbers {
			if board.numbers[j][i].marked {
				markings += 1
				resCol = append(resCol, board.numbers[i][j])
			}
			//fmt.Println(resCol)
			if markings == bingoLength {
				//fmt.Println("COL WINNER")
				return true, resCol
			}
		}
	}

	return false, nil
}

func playBingo(numbers []int, boards []Board) (winningBoard Board, winningNumber int) {
	hasBingo := false

	for !hasBingo {
		for _, n := range numbers {
			//fmt.Println("playing ", n)

			for _, board := range boards {
				for _, row := range board.numbers {
					for i, v := range row {
						if n == v.number {
							row[i].marked = true
						}
					}
				}
				win, _ := checkBingo(board)
				if win {
					hasBingo = true
					//fmt.Println("We have winner! Board ", i)
					//fmt.Println("Winning line", winningLine)
					//fmt.Println("Winning number", n)
					return board, n
				}
			}
		}
	}

	return Board{}, 0
}

func playBingo2(numbers []int, boards []Board) (winningBoard Board, winningNumber int) {
	boardBingosMax := len(boards)
	bingos := 0

	for bingos < boardBingosMax {
		for _, n := range numbers {
			//fmt.Println("playing ", n)

			for i, board := range boards {
				for _, row := range board.numbers {
					for i, v := range row {
						if n == v.number {
							row[i].marked = true
						}
					}
				}
				win, _ := checkBingo(board)
				if win && !boards[i].bingo && bingos < boardBingosMax-1 {
					boards[i].bingo = true
					//fmt.Println("Setting board ", i, "to already won!")
					bingos++
					//fmt.Println("TJA We have false winner! Board ", i)
					//fmt.Println("TJA Winning line", winningLine)
					//fmt.Println("TJA Winning number", n)
				} else if win && !boards[i].bingo && bingos == boardBingosMax-1 {
					bingos++
					boards[i].bingo = true
					//fmt.Println("We have winner! Board ", i)
					//fmt.Println("Winning line", winningLine)
					//fmt.Println("Winning number", n)
					return board, n
				}
				//fmt.Println("BINGOS: ", bingos)
			}
		}
	}

	return Board{}, 0
}

func part1(bingo []string) int {

	numbers, boards := parseBingo(bingo)

	//fmt.Println("Playing with ", numbers)

	winningBoard, winningNumber := playBingo(numbers, boards)

	sum := 0

	for _, row := range winningBoard.numbers {
		for _, n := range row {
			if !n.marked {
				sum += n.number
			}
		}
	}

	// debug
	//for i, bs := range boards {
	//	fmt.Println("Board ", i)
	//	for _, r := range bs.numbers {
	//		fmt.Println(r)
	//	}
	//}

	return sum * winningNumber
}

func part2(bingo []string) int {
	numbers, boards := parseBingo(bingo)

	//fmt.Println("Playing with ", numbers)

	winningBoard, winningNumber := playBingo2(numbers, boards)

	sum := 0

	for _, row := range winningBoard.numbers {
		for _, n := range row {
			if !n.marked {
				sum += n.number
			}
		}
	}

	// debug
	//for i, bs := range boards {
	//	fmt.Println("Board ", i)
	//	for _, r := range bs.numbers {
	//		fmt.Println(r)
	//	}
	//}

	return sum * winningNumber

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
