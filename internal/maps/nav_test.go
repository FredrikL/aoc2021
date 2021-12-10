package maps

import (
	"bufio"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func load_day10() []string {
	file, _ := os.Open("../../test/day10_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Test_ShouldReturnInvalid(t *testing.T) {

	res, _ := ValidateChunks("(]")
	assert.Equal(t, "]", res)

	res, _ = ValidateChunks("{()()()>")
	assert.Equal(t, ">", res)

	res, _ = ValidateChunks("(((()))}")
	assert.Equal(t, "}", res)

	res, _ = ValidateChunks("<([]){()}[{}])")
	assert.Equal(t, ")", res)
}

func Test_RunDay10ExampleData(t *testing.T) {
	data := []string{"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]"}
	invalid := []string{}

	for _, line := range data {
		res, _ := ValidateChunks(line)
		if res != "" {
			invalid = append(invalid, res)
		}
	}

	assert.Equal(t, []string{"}", ")", "]", ")", ">"}, invalid)

	sum := 0
	for _, v := range invalid {
		if v == ")" {
			sum += 3
		}
		if v == "}" {
			sum += 1197
		}
		if v == "]" {
			sum += 57
		}
		if v == ">" {
			sum += 25137
		}
	}
	assert.Equal(t, 26397, sum)
}

func Test_Day10P1(t *testing.T) {
	data := load_day10()
	invalid := []string{}

	for _, line := range data {
		res, _ := ValidateChunks(line)
		if res != "" {
			invalid = append(invalid, res)
		}
	}

	sum := 0
	for _, v := range invalid {
		if v == ")" {
			sum += 3
		}
		if v == "}" {
			sum += 1197
		}
		if v == "]" {
			sum += 57
		}
		if v == ">" {
			sum += 25137
		}
	}
	assert.Equal(t, 389589, sum)
}

func Test_ReturnStackForIncomplete(t *testing.T) {
	res, stack := ValidateChunks("[({(<(())[]>[[{[]{<()<>>")
	assert.Equal(t, "", res)
	stack = ReverseStack(stack)
	assert.Equal(t, []string{"}", "}", "]", "]", ")", "}", ")", "]"}, stack)
	sum := CalcStackScore(stack)
	assert.Equal(t, 288957, sum)

	res, stack = ValidateChunks("[(()[<>])]({[<{<<[]>>(")
	assert.Equal(t, "", res)
	stack = ReverseStack(stack)
	assert.Equal(t, []string{")", "}", ">", "]", "}", ")"}, stack)
	sum = CalcStackScore(stack)
	assert.Equal(t, 5566, sum)
}

func Test_Day10P2(t *testing.T) {
	data := load_day10()
	sums := []int{}

	for _, line := range data {
		res, stack := ValidateChunks(line)
		if res == "" {
			stack = ReverseStack(stack)
			sum := CalcStackScore(stack)
			sums = append(sums, sum)
		}
	}

	sort.Ints(sums)

	pos := len(sums) / 2

	assert.Equal(t, 1190420163, sums[pos])
}
