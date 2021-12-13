package aoc

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput string = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func load_day13() string {
	file, _ := os.Open("../test/day13_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

func Test_D13ShouldParseExampleInput(t *testing.T) {
	pairs, fold := ParseOrigamiInstructions(exampleInput)

	assert.Len(t, pairs, 18)
	assert.Len(t, fold, 2)
}

func Test_ShouldFoldExampleInput(t *testing.T) {
	pairs, fold := ParseOrigamiInstructions(exampleInput)

	pairs = FoldOrigami(pairs, fold[0])
	assert.Len(t, pairs, 17)

	pairs = FoldOrigami(pairs, fold[1])
	assert.Len(t, pairs, 16)
}

func Test_ShouldFoldDay13Input(t *testing.T) {
	input := load_day13()
	pairs, fold := ParseOrigamiInstructions(input)

	pairs = FoldOrigami(pairs, fold[0])
	assert.Len(t, pairs, 807)
}

func Test_ShouldFoldDay13P2(t *testing.T) {
	input := load_day13()
	pairs, fold := ParseOrigamiInstructions(input)

	for _, f := range fold {
		pairs = FoldOrigami(pairs, f)
	}

	assert.Len(t, pairs, 98)

	canvas := make([][]string, 6)
	for i := range canvas {
		canvas[i] = make([]string, 40)
		for j := range canvas[i] {
			canvas[i][j] = "."
		}
	}
	for _, p := range pairs {
		canvas[p.y][p.x] = "#"
	}
	for i := range canvas {
		line := strings.Join(canvas[i], "")
		t.Log(line)
	}
}
