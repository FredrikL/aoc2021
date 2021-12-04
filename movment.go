package aoc

import (
	"strconv"
	"strings"
)

type Direction int8

const (
	Unset Direction = iota
	Forward
	Up
	Down
)

type Command struct {
	action Direction
	units  int
}

func CommandParser(c string) Command {
	parts := strings.Split(c, " ")

	units, _ := strconv.Atoi(parts[1])
	var action Direction

	switch parts[0] {
	case "forward":
		action = Forward
	case "down":
		action = Down
	case "up":
		action = Up
	}

	return Command{
		action: action,
		units:  units,
	}
}

func ApplyCommand(horizontal, depth int, command Command) (int, int) {
	switch command.action {
	case Forward:
		horizontal += command.units
	case Down:
		depth += command.units
	case Up:
		depth -= command.units
	}
	return horizontal, depth
}

func ApplyCommandWithAim(horizontal, depth, aim int, command Command) (int, int, int) {
	switch command.action {
	case Forward:
		horizontal += command.units
		depth = depth + (command.units * aim)
	case Down:
		aim += command.units
	case Up:
		aim -= command.units
	}
	return horizontal, depth, aim
}
