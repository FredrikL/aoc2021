package aoc

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const d5ExampleInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func Test_ShouldConvertRowToLine(t *testing.T) {
	input := "2,2 -> 2,1"
	expected := Line{Start: Point{X: 2, Y: 2}, End: Point{X: 2, Y: 1}}

	ins := ReadRow(input)

	assert.Equal(t, expected, ins)
}

func Test_LoadAndFilterInput(t *testing.T) {
	result := ConvertInputSplit(d5ExampleInput, true)

	assert.Len(t, result, 6)
}

func Test_ExpandLine(t *testing.T) {
	res := ExpandLine(Line{Start: Point{X: 2, Y: 2}, End: Point{X: 2, Y: 1}})

	assert.Equal(t, []Point{{2, 2}, {2, 1}}, res)
}

func Test_CalcVentData(t *testing.T) {
	input := "2,2 -> 2,1"
	loaded := ConvertInputSplit(input, true)
	assert.Len(t, loaded, 1)

	m := ApplyVentData(loaded)
	assert.Equal(t, map[string]int(map[string]int{"2,2": 1, "2,1": 1}), m)
}

func Test_CalcVentDataExampleInput(t *testing.T) {
	lines := ConvertInputSplit(d5ExampleInput, true)
	ventMap := ApplyVentData(lines)

	c := CountLarge(ventMap)
	assert.Equal(t, 5, c)
}

func Test_CalcDay5P1(t *testing.T) {
	lines := ConvertInput(load_day5(), true)
	ventMap := ApplyVentData(lines)

	c := CountLarge(ventMap)
	assert.Equal(t, 3990, c)
}

func Test_ShouldLoadAllData(t *testing.T) {
	result := ConvertInputSplit(d5ExampleInput, false)

	assert.Len(t, result, 10)
}

func Test_ExpandLineDiagonal(t *testing.T) {
	res := ExpandLine(Line{Start: Point{X: 1, Y: 1}, End: Point{X: 3, Y: 3}})
	assert.Equal(t, []Point{{1, 1}, {2, 2}, {3, 3}}, res)

	res = ExpandLine(Line{Start: Point{X: 9, Y: 7}, End: Point{X: 7, Y: 9}})
	assert.Equal(t, []Point{{9, 7}, {8, 8}, {7, 9}}, res)

}

func Test_CalcDay5P2(t *testing.T) {
	lines := ConvertInput(load_day5(), false)
	ventMap := ApplyVentData(lines)

	c := CountLarge(ventMap)
	assert.Equal(t, 21305, c)
}

func load_day5() []string {
	file, _ := os.Open("day5_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
