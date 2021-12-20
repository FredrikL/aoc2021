package maps

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput string = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

func load_day20() string {
	file, _ := os.Open("../../test/day20_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

func Test_ParseScannerExampleInput(t *testing.T) {
	lookup, state := ParseScannerInput(exampleInput)

	assert.Len(t, lookup, 512)
	assert.Len(t, state, 10)
}

func Test_StepScannerState(t *testing.T) {
	lookup, state := ParseScannerInput(exampleInput)

	state = StepScannerState(lookup, state)

	assert.Len(t, state, 24)

	state = StepScannerState(lookup, state)

	assert.Len(t, state, 35)
}

func test_Day20P1(t *testing.T) {
	lookup, state := ParseScannerInput(load_day20())

	assert.Len(t, lookup, 512)
	assert.Len(t, state, 5004)

	state = StepScannerState(lookup, state)

	state = StepScannerState(lookup, state)

	state = Trim(state)

	assert.Equal(t, 5361, len(state))
}

func Test_Day20P2ExampleData(t *testing.T) {
	lookup, state := ParseScannerInput(exampleInput)

	for i := 0; i < 50; i++ {
		state = StepScannerState(lookup, state)

	}

	assert.Equal(t, 3351, len(state))
}

func test_Day20P2(t *testing.T) {
	lookup, state := ParseScannerInput(load_day20())

	for i := 0; i < 50; i++ {
		state = StepScannerState(lookup, state)
		if i%2 == 1 {
			state = Trim(state)
		}
	}

	assert.Equal(t, 16826, len(state))
}
