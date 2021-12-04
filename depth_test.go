package aoc

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExampleInput(t *testing.T) {
	input := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	inc := DepthIncreseCount(input)

	assert.Equal(t, 7, inc)
}

func Test_Day1P1Input(t *testing.T) {
	input := load()

	inc := DepthIncreseCount(input)

	assert.Equal(t, 1167, inc)
}

func Test_MergeMesurmentsBy3(t *testing.T) {
	input := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}
	expected := []int{607,
		618,
		618,
		617,
		647,
		716,
		769,
		792}

	merged := MergeMesurments(input)

	assert.Equal(t, expected, merged)
}

func Test_ExampleInputp2(t *testing.T) {
	input := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	merged := MergeMesurments(input)

	inc := DepthIncreseCount(merged)

	assert.Equal(t, 5, inc)
}

func Test_Day1P2Input(t *testing.T) {
	input := load()
	merged := MergeMesurments(input)

	inc := DepthIncreseCount(merged)

	assert.Equal(t, 1130, inc)
}

func load() []int {
	file, _ := os.Open("day1_input.txt")
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, v)
	}

	return lines
}
