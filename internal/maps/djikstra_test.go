package maps

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleMap string = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func toInputMap(input string) ([][]int, [][]int, [][]bool) {
	result := [][]int{}

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		r := []int{}
		rp := strings.Split(row, "")
		for _, rv := range rp {
			v, _ := strconv.Atoi(rv)
			r = append(r, v)
		}
		result = append(result, r)
	}

	risk := make([][]int, len(result))
	seen := make([][]bool, len(result))
	for i := range risk {
		risk[i] = make([]int, len(result[0]))
		for j := range risk[i] {
			risk[i][j] = 10000
		}
		seen[i] = make([]bool, len(result[0]))
	}

	return result, risk, seen
}

func load_day15() string {
	file, _ := os.Open("../../test/day15_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

func Test_LoadExamplData(t *testing.T) {
	mp, risk, _ := toInputMap(exampleMap)

	assert.Len(t, mp, 10)
	assert.Len(t, mp[1], 10)

	assert.Len(t, risk, 10)
	assert.Len(t, risk[1], 10)
	assert.Equal(t, 0, risk[1][0])
}

func Test_WalkExample(t *testing.T) {
	mp, risk, seen := toInputMap(exampleMap)
	mp[0][0] = 0
	WalkMap(&mp, &risk, &seen)

	assert.Equal(t, 40, risk[9][9])
}
func Test_WalkDay15P1(t *testing.T) {
	mp, risk, seen := toInputMap(load_day15())
	mp[0][0] = 0
	WalkMap(&mp, &risk, &seen)

	assert.Equal(t, 462, risk[len(mp)-1][len(mp[0])-1])
}
