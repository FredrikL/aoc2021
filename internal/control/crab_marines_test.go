package control

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func Test_GetMinMax(t *testing.T) {
	input := conv("16,1,2,0,4,2,7,1,2,14")

	min, max := GetMinMax(input)

	assert.Equal(t, 0, min)
	assert.Equal(t, 16, max)
}*/

func Test_CalcFuel(t *testing.T) {
	input := conv("16,1,2,0,4,2,7,1,2,14")

	use := CalcFuel(input, 2)

	assert.Equal(t, 206, use)

	use = CalcFuel(input, 5)

	assert.Equal(t, 168, use)

}

func Test_FindBestPos(t *testing.T) {
	input := conv("16,1,2,0,4,2,7,1,2,14")

	min, max := GetMinMax(input)

	use := FindBestPos(input, min, max)

	assert.Equal(t, 168, use)
}

/*

func Test_Day7P1(t *testing.T) {
	input := load_day7()
	min, max := GetMinMax(input)

	use := FindBestPos(input, min, max)

	assert.Equal(t, 37, use)
}
*/
func Test_Day7P2(t *testing.T) {
	input := load_day7()
	min, max := GetMinMax(input)

	use := FindBestPos(input, min, max)

	assert.Equal(t, 100727924, use)
}

func conv(input string) []int {
	parts := strings.Split(input, ",")
	var res []int
	for _, v := range parts {
		i, _ := strconv.Atoi(v)
		res = append(res, i)
	}
	return res
}

func load_day7() []int {
	file, _ := os.Open("../../test/day7_input.txt")
	defer file.Close()
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	return conv(line)
}
