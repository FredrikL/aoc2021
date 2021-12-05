package aoc

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string = []string{"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010"}

func Test_GetMostCommonForColumn(t *testing.T) {
	result := GetMostCommonForColumn(input, 0)

	assert.Equal(t, "1", result)

	result = GetMostCommonForColumn(input, 1)

	assert.Equal(t, "0", result)
	result = GetMostCommonForColumn(input, 2)

	assert.Equal(t, "1", result)
	result = GetMostCommonForColumn(input, 3)

	assert.Equal(t, "1", result)
	result = GetMostCommonForColumn(input, 4)

	assert.Equal(t, "0", result)
}

func Test_GetGammaRate(t *testing.T) {
	result := GetGammaRate(input)

	assert.Equal(t, int64(22), result)
}

func Test_GetEpsilonRate(t *testing.T) {
	result := GetEpsilonRate(input)

	assert.Equal(t, int64(9), result)
}

func Test_GetExampleInputResult(t *testing.T) {
	g, e := GetGammaRate(input), GetEpsilonRate(input)

	assert.Equal(t, int64(198), g*e)
}

func Test_Day3P1(t *testing.T) {
	data := load_day3()
	g, e := GetGammaRate(data), GetEpsilonRate(data)

	assert.Equal(t, int64(738234), g*e)
}

func Test_GetOxygenRating(t *testing.T) {
	res := GetOxygenRating(input)

	assert.Equal(t, int64(23), res)
}

func Test_GetCo2Rating(t *testing.T) {
	res := GetCo2Rating(input)

	assert.Equal(t, int64(10), res)
}

func Test_Day3P2ExampleInput(t *testing.T) {
	o, co2 := GetOxygenRating(input), GetCo2Rating(input)

	assert.Equal(t, int64(230), o*co2)
}

func Test_Day3P3(t *testing.T) {
	data := load_day3()
	o, co2 := GetOxygenRating(data), GetCo2Rating(data)

	assert.Equal(t, int64(3969126), o*co2)
}

func load_day3() []string {
	file, _ := os.Open("../test/day3_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
