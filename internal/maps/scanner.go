package maps

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

func bounds(current map[pos]bool) (int, int, int, int) {
	xmin, xmax, ymin, ymax := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt

	for k := range current {
		if k.x < xmin {
			xmin = k.x
		}
		if k.x > xmax {
			xmax = k.x
		}
		if k.y < ymin {
			ymin = k.y
		}
		if k.y > ymax {
			ymax = k.y
		}
	}

	return xmin, xmax, ymin, ymax
}

func ParseScannerInput(input string) (string, map[pos]bool) {
	parts := strings.Split(input, "\n")

	res := make(map[pos]bool)

	for i := 2; i < len(parts); i++ {
		row := strings.Split(parts[i], "")
		for k, v := range row {
			if v == "#" {
				key := pos{x: i, y: k}
				res[key] = true
			}
		}
	}
	return parts[0], res
}

func StepScannerState(lookup string, current map[pos]bool) map[pos]bool {
	xmin, xmax, ymin, ymax := bounds(current)

	res := make(map[pos]bool)

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {

			for i := (x - 10); i <= (x + 10); i++ {
				for j := (y - 10); j <= (y + 10); j++ {
					v := GetScannerValueAtPos(lookup, current, i, j)
					if v {
						k := pos{x: i, y: j}
						res[k] = true
					}
				}
			}
		}
	}

	return res
}

func Trim(current map[pos]bool) map[pos]bool {
	minX, maxX, minY, maxY := bounds(current)

	res := make(map[pos]bool)

	borderThickness := 0

	middleY := (minY + maxY) / 2
	for x := minX; x < maxX; x++ {
		k := pos{x: x, y: middleY}
		if !current[k] {
			borderThickness = x - minX
			break
		}
	}

	for y := minY + borderThickness; y <= maxY-borderThickness; y++ {
		for x := minX + borderThickness; x <= maxX-borderThickness; x++ {
			k := pos{x: x, y: y}
			if current[k] {
				res[k] = true
			}
		}
	}

	return res
}

func GetScannerValueAtPos(lookup string, current map[pos]bool, x, y int) bool {
	key := make([]string, 9)

	idx := 0
	for i := (x - 1); i <= (x + 1); i++ {
		for j := (y - 1); j <= (y + 1); j++ {
			k := pos{x: i, y: j}
			if current[k] {
				key[idx] = "1"
			} else {
				key[idx] = "0"
			}
			idx++
		}
	}

	j := strings.Join(key, "")
	lookupIndex, _ := strconv.ParseInt(j, 2, 64)

	return lookup[int(lookupIndex):int(lookupIndex)+1] == "#"
}

func ShowState(state map[string]bool) {
	rows := make([][]string, 110)
	for i := range rows {
		rows[i] = make([]string, 110)
		for j := range rows[i] {
			rows[i][j] = "."
		}
	}

	for key, v := range state {
		pos := strings.Split(key, ":")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])

		if v {
			rows[x][y] = "#"
		}
	}

	for _, r := range rows {
		row := strings.Join(r, "")
		fmt.Println(row)
	}
}
