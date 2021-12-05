package aoc

import (
	"strconv"
	"strings"
)

type Cell struct {
	number int
	drawn  bool
}

type Board = [5][5]Cell

func CreateBoard(lines []string) Board {
	res := Board{}

	splitFn := func(c rune) bool {
		return c == ' '
	}

	for i := 0; i < len(lines); i++ {
		parts := strings.FieldsFunc(strings.TrimSpace(lines[i]), splitFn)
		for j := 0; j < len(parts); j++ {
			v, _ := strconv.Atoi(strings.TrimSpace(parts[j]))

			res[i][j] = Cell{number: v, drawn: false}
		}

	}
	return res
}

func LoadInput(lines []string) ([]int, []Board) {
	nums := strings.Split(lines[0], ",")
	var numbers []int
	for i := 0; i < len(nums); i++ {
		v, _ := strconv.Atoi(nums[i])
		numbers = append(numbers, v)
	}
	var board []Board
	for i := 0; i < ((len(lines) - 1) / 6); i++ {
		index := i*6 + 1
		s := lines[(index + 1):(index + 6)]
		b := CreateBoard(s)
		board = append(board, b)
	}

	return numbers, board
}

func MarkBoard(board Board, number int) Board {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j].number == number {
				board[i][j].drawn = true
				return board
			}
		}
	}
	return board
}

func HasWinningRow(board Board) bool {
	for i := 0; i < 5; i++ {
		if board[i][0].drawn && board[i][1].drawn && board[i][2].drawn && board[i][3].drawn && board[i][4].drawn {
			return true
		}
	}
	return false
}

func HasWinningColumn(board Board) bool {
	for i := 0; i < 5; i++ {
		if board[0][i].drawn && board[1][i].drawn && board[2][i].drawn && board[3][i].drawn && board[4][i].drawn {
			return true
		}
	}
	return false
}

func Play(numbers []int, boards []Board) (int, Board) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(boards); j++ {
			boards[j] = MarkBoard(boards[j], numbers[i])
		}
		for j := 0; j < len(boards); j++ {
			if HasWinningRow(boards[j]) || HasWinningColumn(boards[j]) {
				return numbers[i], boards[j]
			}
		}
	}

	return 0, Board{}
}

func PlayToEnd(numbers []int, boards []Board) (int, Board) {
	for i := 0; i < len(numbers); i++ {
		win := make([]bool, len(boards))
		numLoose := 0
		for j := 0; j < len(boards); j++ {
			if HasWinningRow(boards[j]) || HasWinningColumn(boards[j]) {
				win[j] = true
			} else {
				numLoose++
			}
		}

		for j := 0; j < len(boards); j++ {
			boards[j] = MarkBoard(boards[j], numbers[i])
		}

		for j := 0; j < len(boards); j++ {
			if (HasWinningRow(boards[j]) || HasWinningColumn(boards[j])) &&
				numLoose == 1 &&
				!win[j] {
				return numbers[i], boards[j]
			}
		}
	}

	return 0, Board{}
}

func CalcWinValue(number int, board Board) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[i][j].drawn {
				sum += board[i][j].number
			}
		}
	}

	return sum * number
}
