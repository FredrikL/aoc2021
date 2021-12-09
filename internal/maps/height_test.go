package maps

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func toInput(str []string) [][]int {
	ret := make([][]int, len(str))

	for i, v := range str {
		parts := strings.Split(v, "")
		for _, p := range parts {
			input, _ := strconv.Atoi(p)
			ret[i] = append(ret[i], input)
		}
	}

	return ret
}

func load_day9() []string {
	file, _ := os.Open("../../test/day9_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Test_FindLowPoints(t *testing.T) {
	x := []string{"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	input := toInput(x)

	result := ReturnLowPoints(input)

	// assert.Equal(t, []int{1, 0, 5, 5}, result)

	r := []int{}
	for _, v := range result {
		r = append(r, v.value)
	}

	sum := SumResult(r)

	assert.Equal(t, 15, sum)
}

func Test_Day9P1(t *testing.T) {
	input := toInput(load_day9())

	result := ReturnLowPoints(input)
	r := []int{}
	for _, v := range result {
		r = append(r, v.value)
	}
	sum := SumResult(r)

	assert.Equal(t, 462, sum)
}

func Test_FindBasinSize(t *testing.T) {
	x := []string{"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	input := toInput(x)

	result := ReturnLowPoints(input)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != 9 {
				input[i][j] = 0
			}
		}
	}

	start := result[0]
	size := FindBasinSize(input, start.x, start.y)
	assert.Equal(t, 3, size)

	start = result[1]
	size = FindBasinSize(input, start.x, start.y)
	assert.Equal(t, 9, size)

	start = result[2]
	size = FindBasinSize(input, start.x, start.y)
	assert.Equal(t, 14, size)

	start = result[3]
	size = FindBasinSize(input, start.x, start.y)
	assert.Equal(t, 9, size)
}

func Test_FindSumOfThreeLargest(t *testing.T) {
	x := []string{"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	input := toInput(x)

	result := ReturnLowPoints(input)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != 9 {
				input[i][j] = 0
			}
		}
	}
	results := []int{}
	for _, r := range result {
		size := FindBasinSize(input, r.x, r.y)
		results = append(results, size)
	}

	sort.Ints(results)

	l := len(results)
	sum := results[l-1] * results[l-2] * results[l-3]

	assert.Equal(t, 1134, sum)
}

func Test_Day9P2(t *testing.T) {
	input := toInput(load_day9())

	result := ReturnLowPoints(input)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != 9 {
				input[i][j] = 0
			}
		}
	}
	results := []int{}
	for _, r := range result {
		size := FindBasinSize(input, r.x, r.y)
		results = append(results, size)
	}

	sort.Ints(results)

	l := len(results)
	sum := results[l-1] * results[l-2] * results[l-3]

	assert.Equal(t, 1397760, sum)
}
