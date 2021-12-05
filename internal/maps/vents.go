package maps

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func toPoint(pair string) Point {
	parts := strings.Split(pair, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Point{X: x, Y: y}
}

func ReadRow(row string) Line {
	ret := Line{}
	parts := strings.Split(row, " -> ")

	ret.Start = toPoint(parts[0])
	ret.End = toPoint(parts[1])

	return ret
}

func ConvertInputSplit(input string, filter bool) []Line {
	rows := strings.Split(input, "\n")
	return ConvertInput(rows, filter)
}

func ConvertInput(rows []string, filter bool) []Line {
	var result []Line
	for _, v := range rows {
		ins := ReadRow(v)
		if filter {
			if (ins.Start.X == ins.End.X) || ins.Start.Y == ins.End.Y {
				result = append(result, ins)
			}
			continue
		} else {
			result = append(result, ins)
		}
	}
	return result
}

func getRange(a, b int) []int {
	var ret []int
	if a < b {
		for i := a; i <= b; i++ {
			ret = append(ret, i)
		}
	} else {
		for i := a; i >= b; i-- {
			ret = append(ret, i)
		}
	}
	return ret
}

func ExpandLine(line Line) []Point {
	var ret []Point
	if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
		if line.Start.X == line.End.X {
			if line.Start.Y < line.End.Y {
				for i := line.Start.Y; i <= line.End.Y; i++ {
					ret = append(ret, Point{line.Start.X, i})
				}
			} else {
				for i := line.Start.Y; i >= line.End.Y; i-- {
					ret = append(ret, Point{line.Start.X, i})
				}
			}
		} else {
			if line.Start.X < line.End.X {
				for i := line.Start.X; i <= line.End.X; i++ {
					ret = append(ret, Point{i, line.Start.Y})
				}
			} else {
				for i := line.Start.X; i >= line.End.X; i-- {
					ret = append(ret, Point{i, line.Start.Y})
				}
			}
		}
	} else {
		xp := getRange(line.Start.X, line.End.X)
		yp := getRange(line.Start.Y, line.End.Y)
		for i, _ := range xp {
			ret = append(ret, Point{xp[i], yp[i]})
		}
	}

	return ret
}

func ApplyVentData(lines []Line) map[string]int {
	m := make(map[string]int)

	for _, line := range lines {
		points := ExpandLine(line)
		for _, point := range points {
			pos := fmt.Sprintf("%d,%d", point.X, point.Y)
			m[pos]++
		}
	}

	return m
}

func CountLarge(vents map[string]int) int {
	sum := 0
	for _, v := range vents {
		if v > 1 {
			sum++
		}
	}
	return sum
}
