package aoc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLanterns(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	res := CreateLanterns(input)

	assert.Equal(t, map[int]int(map[int]int{1: 1, 2: 1, 3: 2, 4: 1}), res)
}

func Test_StepLanterns(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	res := CreateLanterns(input)
	assert.Equal(t, map[int]int(map[int]int{1: 1, 2: 1, 3: 2, 4: 1}), res)

	res = StepLanterns(res)
	assert.Equal(t, map[int]int(map[int]int{0: 1, 1: 1, 2: 2, 3: 1}), res)

	res = StepLanterns(res)
	// 1, 2, 1, 6, 0, 8
	assert.Equal(t, map[int]int(map[int]int{0: 1, 1: 2, 2: 1, 6: 1, 8: 1}), res)

	res = StepLanterns(res)
	// 0, 1, 0, 5, 6, 7, 8
	assert.Equal(t, map[int]int(map[int]int{0: 2, 1: 1, 5: 1, 6: 1, 7: 1, 8: 1}), res)

	res = StepLanterns(res)
	// 6, 0, 6, 4, 5, 6, 7, 8, 8
	assert.Equal(t, map[int]int(map[int]int{0: 1, 4: 1, 5: 1, 6: 3, 7: 1, 8: 2}), res)
}

func Test_Step18Generations(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	res := CreateLanterns(input)
	for i := 0; i < 18; i++ {
		res = StepLanterns(res)
	}

	count := 0
	for _, v := range res {
		count += v
	}

	assert.Equal(t, 26, count)

}

func Test_StepExampleData80Generations(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	res := CreateLanterns(input)
	for i := 0; i < 80; i++ {
		res = StepLanterns(res)
	}

	count := 0
	for _, v := range res {
		count += v
	}

	assert.Equal(t, 5934, count)
}

func Test_Day6P1(t *testing.T) {
	input := load_day6()
	res := CreateLanterns(input)
	for i := 0; i < 80; i++ {
		res = StepLanterns(res)
	}

	count := 0
	for _, v := range res {
		count += v
	}

	assert.Equal(t, 376194, count)

}

func Test_StepExampleData256Generations(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	res := CreateLanterns(input)
	for i := 0; i < 256; i++ {
		res = StepLanterns(res)
	}

	count := 0
	for _, v := range res {
		count += v
	}

	assert.Equal(t, 26984457539, count)
}

func Test_Day6P2(t *testing.T) {
	input := load_day6()
	res := CreateLanterns(input)
	for i := 0; i < 256; i++ {
		res = StepLanterns(res)
	}

	count := 0
	for _, v := range res {
		count += v
	}

	assert.Equal(t, 1693022481538, count)

}

func load_day6() []int {
	file, _ := os.Open("../test/day6_input.txt")
	defer file.Close()
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}
	var res []int
	parts := strings.Split(line, ",")
	for _, v := range parts {
		i, _ := strconv.Atoi(v)
		res = append(res, i)
	}

	return res
}
