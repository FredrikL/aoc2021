package aoc

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountDigitsInRow(t *testing.T) {
	input := "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"
	count := CountDigitsInRow(input)
	assert.Equal(t, 2, count)
}

func Test_CountExampleData(t *testing.T) {
	input := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
	count := 0
	for _, v := range input {
		count += CountDigitsInRow(v)
	}
	assert.Equal(t, 26, count)
}

func Test_Day8P1(t *testing.T) {
	input := load_day8()
	count := 0
	for _, v := range input {
		count += CountDigitsInRow(v)
	}
	assert.Equal(t, 278, count)
}

func Test_WithLookup(t *testing.T) {
	res := MapDigits("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")

	assert.Equal(t, 1, res["ab"])
	assert.Equal(t, 8, res["acedgfb"])
	assert.Equal(t, 4, res["eafb"])
	assert.Equal(t, 7, res["dab"])
	assert.Equal(t, 0, res["cagedb"])
	assert.Equal(t, 9, res["cefabd"])
	assert.Equal(t, 6, res["cdfgeb"])
	assert.Equal(t, 2, res["gcdfa"])
	assert.Equal(t, 3, res["fbcad"])
	assert.Equal(t, 5, res["cdfbe"])

	value := MapRow("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.Equal(t, 5353, value)

	value = MapRow("edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc")
	assert.Equal(t, 9781, value)
}

func Test_CountExampleDataP2(t *testing.T) {
	input := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
	count := 0
	for _, v := range input {
		count += MapRow(v)
	}
	assert.Equal(t, 61229, count)
}

func Test_Day8P2(t *testing.T) {
	input := load_day8()
	count := 0
	for _, v := range input {
		count += MapRow(v)
	}
	assert.Equal(t, 986179, count)
}

func load_day8() []string {
	file, _ := os.Open("../test/day8_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
