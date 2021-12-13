package aoc

import (
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

type Fold struct {
	edge string
	pos  int
}

func ParseOrigamiInstructions(input string) ([]Pair, []Fold) {
	lines := strings.Split(input, "\n")
	pairs := []Pair{}
	fold := []Fold{}
	pairsDone := false
	for _, line := range lines {
		if line == "" {
			pairsDone = true
			continue
		}
		if !pairsDone {
			parts := strings.Split(line, ",")
			pair := Pair{}
			v, _ := strconv.Atoi(parts[0])
			pair.x = v
			v, _ = strconv.Atoi(parts[1])
			pair.y = v
			pairs = append(pairs, pair)
		} else {
			line = strings.Replace(line, "fold along ", "", 1)
			parts := strings.Split(line, "=")
			f := Fold{}
			v, _ := strconv.Atoi(parts[1])
			f.pos = v
			f.edge = parts[0]
			fold = append(fold, f)
		}
	}
	return pairs, fold
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FoldOrigami(pairs []Pair, instruction Fold) []Pair {
	updated := []Pair{}
	for _, p := range pairs {
		if instruction.edge == "y" {
			if p.y < instruction.pos {
				updated = append(updated, p)
			} else if p.y > instruction.pos {
				new_y := Abs((p.y - instruction.pos) - instruction.pos)
				p.y = new_y
				updated = append(updated, p)
			}
		} else {
			if p.x < instruction.pos {
				updated = append(updated, p)
			} else if p.x > instruction.pos {
				new_x := Abs((p.x - instruction.pos) - instruction.pos)
				p.x = new_x
				updated = append(updated, p)
			}
		}
	}

	res := []Pair{}
	// clear duplictes
	for _, p := range updated {
		found := false
		for _, v := range res {
			if p.x == v.x && p.y == v.y {
				found = true
			}
		}
		if !found {
			res = append(res, p)
		}
	}

	return res
}
