package aoc

import (
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

func toRange(str string) (int, int) {
	p := strings.Split(str, "=")
	lh := strings.Split(p[1], "..")
	l, _ := strconv.Atoi(lh[0])
	h, _ := strconv.Atoi(lh[1])

	if l < -50 {
		l = -50
	}
	if h > 50 {
		h = 50
	}
	return l, h
}

func ToCubes(input string) map[Cube]bool {
	res := make(map[Cube]bool)

	p := strings.Split(input, " ")
	ranges := strings.Split(p[1], ",")
	xl, xh := toRange(ranges[0])
	yl, yh := toRange(ranges[1])
	zl, zh := toRange(ranges[2])

	for x := xl; x <= xh; x++ {
		for y := yl; y <= yh; y++ {
			for z := zl; z <= zh; z++ {
				res[Cube{x, y, z}] = true
			}
		}
	}

	return res
}

func Merge(a, b map[Cube]bool) map[Cube]bool {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func Remove(a, b map[Cube]bool) map[Cube]bool {
	for k := range b {
		delete(a, k)
	}
	return a
}
