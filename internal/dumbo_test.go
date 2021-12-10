package aoc

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func handleLine(line string) []int {
	parts := strings.Split(line, "")
	res := []int{}
	for _, d := range parts {
		v, _ := strconv.Atoi(d)

		res = append(res, v)
	}
	return res
}

func toArray(str string) [][]int {
	lines := strings.Split(str, "\n")
	res := [][]int{}
	for _, l := range lines {
		h := handleLine(l)
		res = append(res, h)
	}
	return res
}

func Test_RunSmallInput(t *testing.T) {
	input := toArray(`11111
19991
19191
19991
11111`)

	//expected := [][]int{[]int{3, 4, 5, 4, 3}, []int{4, 0, 0, 0, 4}, []int{5, 0, 0, 0, 5}, []int{4, 0, 0, 0, 4}, []int{3, 4, 5, 4, 3}}
	state, flashSum := StepDumbo(input)
	//assert.Equal(t, expected, state)
	assert.Equal(t, 9, flashSum)

	//expected = [][]int{[]int{4, 5, 6, 5, 4}, []int{5, 1, 1, 1, 5}, []int{6, 1, 1, 1, 6}, []int{5, 1, 1, 1, 5}, []int{4, 5, 6, 5, 4}}
	_, flashSum = StepDumbo(state)
	//	assert.Equal(t, expected, state)
	assert.Equal(t, 0, flashSum)
}

func Test_RunLargerExample(t *testing.T) {
	state := toArray(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)

	flashSum := 0
	sum := 0
	for i := 0; i < 100; i++ {
		state, sum = StepDumbo(state)
		flashSum += sum
	}
	assert.Equal(t, 1656, flashSum)
}

func Test_Day11P1(t *testing.T) {
	state := toArray(`7313511551
3724855867
2374331571
4438213437
6511566287
6727245532
3736868662
2348138263
2417483121
8812617112`)

	flashSum := 0
	sum := 0
	for i := 0; i < 100; i++ {
		state, sum = StepDumbo(state)
		flashSum += sum
	}
	assert.Equal(t, 1655, flashSum)
}

func Test_FindFullFlash(t *testing.T) {
	state := toArray(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)

	sum := 0
	round := 0
	for ok := true; ok; {
		round++
		state, sum = StepDumbo(state)
		if sum == 100 {
			break
		}
	}
	assert.Equal(t, 195, round)
}

func Test_Day11P2(t *testing.T) {
	state := toArray(`7313511551
3724855867
2374331571
4438213437
6511566287
6727245532
3736868662
2348138263
2417483121
8812617112`)

	sum := 0
	round := 0
	for ok := true; ok; {
		round++
		state, sum = StepDumbo(state)
		if sum == 100 {
			break
		}
	}
	assert.Equal(t, 337, round)
}
