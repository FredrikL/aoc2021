package aoc

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseStringToCommand(t *testing.T) {
	cmd := CommandParser("forward 5")

	assert.Equal(t, Forward, cmd.action)
	assert.Equal(t, 5, cmd.units)

	cmd = CommandParser("up 3")

	assert.Equal(t, Up, cmd.action)
	assert.Equal(t, 3, cmd.units)

	cmd = CommandParser("down 8")

	assert.Equal(t, Down, cmd.action)
	assert.Equal(t, 8, cmd.units)
}

func Test_ApplyCommand(t *testing.T) {
	h, d := ApplyCommand(0, 0, Command{action: Forward, units: 5})

	assert.Equal(t, 5, h)
	assert.Equal(t, 0, d)

	h, d = ApplyCommand(0, 0, Command{action: Down, units: 5})

	assert.Equal(t, 0, h)
	assert.Equal(t, 5, d)

	h, d = ApplyCommand(0, 10, Command{action: Up, units: 5})

	assert.Equal(t, 0, h)
	assert.Equal(t, 5, d)
}

func Test_RunExampleInput(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`
	lines := strings.Split(input, "\n")
	var commands []Command
	for i := 0; i < len(lines); i++ {
		commands = append(commands, CommandParser(lines[i]))
	}

	h, d := 0, 0
	for i := 0; i < len(lines); i++ {
		h, d = ApplyCommand(h, d, commands[i])
	}

	assert.Equal(t, 15, h)
	assert.Equal(t, 10, d)
}

func Test_RunDay2P1Input(t *testing.T) {
	commands := loadD2()

	h, d := 0, 0
	for i := 0; i < len(commands); i++ {
		h, d = ApplyCommand(h, d, commands[i])
	}

	assert.Equal(t, 1996, h)
	assert.Equal(t, 1022, d)
	assert.Equal(t, 2039912, h*d)
}

func Test_ApplyCommandWithAim(t *testing.T) {
	h, d, a := ApplyCommandWithAim(0, 0, 0, Command{action: Forward, units: 5})

	assert.Equal(t, 5, h)
	assert.Equal(t, 0, d)
	assert.Equal(t, 0, a)

	h, d, a = ApplyCommandWithAim(5, 0, 0, Command{action: Down, units: 5})

	assert.Equal(t, 5, h)
	assert.Equal(t, 0, d)
	assert.Equal(t, 5, a)

	h, d, a = ApplyCommandWithAim(5, 0, 5, Command{action: Forward, units: 8})

	assert.Equal(t, 13, h)
	assert.Equal(t, 40, d)
	assert.Equal(t, 5, a)

	h, d, a = ApplyCommandWithAim(13, 40, 5, Command{action: Up, units: 3})

	assert.Equal(t, 13, h)
	assert.Equal(t, 40, d)
	assert.Equal(t, 2, a)

	h, d, a = ApplyCommandWithAim(13, 40, 2, Command{action: Down, units: 8})

	assert.Equal(t, 13, h)
	assert.Equal(t, 40, d)
	assert.Equal(t, 10, a)

	h, d, a = ApplyCommandWithAim(13, 40, 10, Command{action: Forward, units: 2})

	assert.Equal(t, 15, h)
	assert.Equal(t, 60, d)
	assert.Equal(t, 10, a)
}

func Test_RunDay2P2Input(t *testing.T) {
	commands := loadD2()

	h, d, a := 0, 0, 0
	for i := 0; i < len(commands); i++ {
		h, d, a = ApplyCommandWithAim(h, d, a, commands[i])
	}

	assert.Equal(t, 1996, h)
	assert.Equal(t, 972980, d)
	assert.Equal(t, 1942068080, h*d)
}

func loadD2() []Command {
	file, _ := os.Open("day2_input.txt")
	defer file.Close()
	var lines []Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := CommandParser(scanner.Text())
		lines = append(lines, c)
	}

	return lines
}
