package aoc

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateBoard(t *testing.T) {
	input := []string{"22 13 17 11  0",
		"8  2 23  4 24",
		"21  9 14 16  7",
		"6 10  3 18  5",
		"1 12 20 15 19"}

	expected := [5][5]Cell{
		[5]Cell{{number: 22, drawn: false}, {number: 13, drawn: false}, {number: 17, drawn: false}, {number: 11, drawn: false}, {number: 0, drawn: false}},
		[5]Cell{{number: 8, drawn: false}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: false}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: false}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: false}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}

	result := CreateBoard(input)

	assert.Equal(t, expected, result)
}

func Test_LoadInput(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

	22 13 17 11  0
	 8  2 23  4 24
	21  9 14 16  7
	 6 10  3 18  5
	 1 12 20 15 19
	
	 3 15  0  2 22
	 9 18 13 17  5
	19  8  7 25 23
	20 11 10 24  4
	14 21 16 12  6
	
	14 21 17 24  4
	10 16 15  9 19
	18  8 23 26 20
	22 11 13  6  5
	 2  0 12  3  7`
	lines := strings.Split(input, "\n")
	expected := [5][5]Cell{
		[5]Cell{Cell{number: 14, drawn: false}, Cell{number: 21, drawn: false}, Cell{number: 17, drawn: false}, Cell{number: 24, drawn: false}, Cell{number: 4, drawn: false}},
		[5]Cell{Cell{number: 10, drawn: false}, Cell{number: 16, drawn: false}, Cell{number: 15, drawn: false}, Cell{number: 9, drawn: false}, Cell{number: 19, drawn: false}},
		[5]Cell{Cell{number: 18, drawn: false}, Cell{number: 8, drawn: false}, Cell{number: 23, drawn: false}, Cell{number: 26, drawn: false}, Cell{number: 20, drawn: false}},
		[5]Cell{Cell{number: 22, drawn: false}, Cell{number: 11, drawn: false}, Cell{number: 13, drawn: false}, Cell{number: 6, drawn: false}, Cell{number: 5, drawn: false}},
		[5]Cell{Cell{number: 2, drawn: false}, Cell{number: 0, drawn: false}, Cell{number: 12, drawn: false}, Cell{number: 3, drawn: false}, Cell{number: 7, drawn: false}}}

	numbers, boards := LoadInput(lines)

	assert.Equal(t, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}, numbers)
	assert.Equal(t, 3, len(boards))
	assert.Equal(t, expected, boards[2])
}

func Test_MarkBoard(t *testing.T) {
	input := [5][5]Cell{
		[5]Cell{{number: 22, drawn: false}, {number: 13, drawn: false}, {number: 17, drawn: false}, {number: 11, drawn: false}, {number: 0, drawn: false}},
		[5]Cell{{number: 8, drawn: false}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: false}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: false}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: false}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}
	expected := [5][5]Cell{
		[5]Cell{{number: 22, drawn: false}, {number: 13, drawn: false}, {number: 17, drawn: false}, {number: 11, drawn: true}, {number: 0, drawn: false}},
		[5]Cell{{number: 8, drawn: false}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: false}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: false}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: false}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}

	result := MarkBoard(input, 11)

	assert.Equal(t, result, expected)
}

func Test_HasWinningRow(t *testing.T) {
	loose := [5][5]Cell{
		[5]Cell{{number: 22, drawn: false}, {number: 13, drawn: false}, {number: 17, drawn: false}, {number: 11, drawn: false}, {number: 0, drawn: false}},
		[5]Cell{{number: 8, drawn: false}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: false}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: false}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: false}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}
	win := [5][5]Cell{
		[5]Cell{{number: 22, drawn: true}, {number: 13, drawn: true}, {number: 17, drawn: true}, {number: 11, drawn: true}, {number: 0, drawn: true}},
		[5]Cell{{number: 8, drawn: true}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: false}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: false}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: false}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}

	assert.False(t, HasWinningRow(loose))
	assert.True(t, HasWinningRow(win))
}

func Test_HasWinningColumn(t *testing.T) {
	loose := [5][5]Cell{
		[5]Cell{{number: 22, drawn: false}, {number: 13, drawn: false}, {number: 17, drawn: false}, {number: 11, drawn: false}, {number: 0, drawn: false}},
		[5]Cell{{number: 8, drawn: false}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: false}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: false}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: false}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}
	win := [5][5]Cell{
		[5]Cell{{number: 22, drawn: true}, {number: 13, drawn: false}, {number: 17, drawn: true}, {number: 11, drawn: true}, {number: 0, drawn: true}},
		[5]Cell{{number: 8, drawn: true}, {number: 2, drawn: false}, {number: 23, drawn: false}, {number: 4, drawn: false}, {number: 24, drawn: false}},
		[5]Cell{{number: 21, drawn: true}, {number: 9, drawn: false}, {number: 14, drawn: false}, {number: 16, drawn: false}, {number: 7, drawn: false}},
		[5]Cell{{number: 6, drawn: true}, {number: 10, drawn: false}, {number: 3, drawn: false}, {number: 18, drawn: false}, {number: 5, drawn: false}},
		[5]Cell{{number: 1, drawn: true}, {number: 12, drawn: false}, {number: 20, drawn: false}, {number: 15, drawn: false}, {number: 19, drawn: false}}}

	assert.False(t, HasWinningColumn(loose))
	assert.True(t, HasWinningColumn(win))
}

func Test_Play(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

	22 13 17 11  0
	 8  2 23  4 24
	21  9 14 16  7
	 6 10  3 18  5
	 1 12 20 15 19
	
	 3 15  0  2 22
	 9 18 13 17  5
	19  8  7 25 23
	20 11 10 24  4
	14 21 16 12  6
	
	14 21 17 24  4
	10 16 15  9 19
	18  8 23 26 20
	22 11 13  6  5
	 2  0 12  3  7`
	lines := strings.Split(input, "\n")

	numbers, boards := LoadInput(lines)

	winNumber, board := Play(numbers, boards)

	assert.Equal(t, 24, winNumber)

	assert.Equal(t, 4512, CalcWinValue(winNumber, board))
}

func Test_PlayToEnd(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

	22 13 17 11  0
	 8  2 23  4 24
	21  9 14 16  7
	 6 10  3 18  5
	 1 12 20 15 19
	
	 3 15  0  2 22
	 9 18 13 17  5
	19  8  7 25 23
	20 11 10 24  4
	14 21 16 12  6
	
	14 21 17 24  4
	10 16 15  9 19
	18  8 23 26 20
	22 11 13  6  5
	 2  0 12  3  7`
	lines := strings.Split(input, "\n")

	numbers, boards := LoadInput(lines)

	winNumber, board := PlayToEnd(numbers, boards)

	assert.Equal(t, 13, winNumber)

	assert.Equal(t, 1924, CalcWinValue(winNumber, board))
}

func Test_RunDay4P1(t *testing.T) {
	lines := load_day4()

	numbers, boards := LoadInput(lines)

	winNumber, board := Play(numbers, boards)

	assert.Equal(t, 77, winNumber)

	assert.Equal(t, 41503, CalcWinValue(winNumber, board))
}

func Test_RunDay4P2(t *testing.T) {
	lines := load_day4()

	numbers, boards := LoadInput(lines)

	winNumber, board := PlayToEnd(numbers, boards)

	assert.Equal(t, 14, winNumber)

	assert.Equal(t, 3178, CalcWinValue(winNumber, board))
}

func load_day4() []string {
	file, _ := os.Open("day4_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
