package aoc

import (
	"strconv"
	"strings"
)

type Target struct {
	x0, y0, x1, y1 int
}

type Velocity struct {
	x, y int
}

type Position struct {
	x, y int
	maxY int
}

func ParseTargetArea(input string) Target {
	res := Target{}
	parts := strings.Split(input, "=")

	xparts := strings.Split(parts[1], "..")

	v, _ := strconv.Atoi(xparts[0])
	res.x0 = v

	x2 := strings.Split(xparts[1], ",")
	v, _ = strconv.Atoi(x2[0])
	res.x1 = v

	yparts := strings.Split(parts[2], "..")

	v, _ = strconv.Atoi(yparts[0])
	res.y1 = v

	v, _ = strconv.Atoi(yparts[1])
	res.y0 = v

	return res
}

func UpdateVelocity(velocity Velocity) Velocity {
	newvelocity := Velocity{}
	newvelocity.y = velocity.y - 1

	if velocity.x > 0 {
		newvelocity.x = velocity.x - 1
	}
	if velocity.x < 0 {
		newvelocity.x = velocity.x + 1
	}

	return newvelocity
}

func ApplyVelocity(v Velocity, p Position) Position {
	newPos := Position{}
	newPos.x = p.x + v.x
	newPos.y = p.y + v.y
	if newPos.y > p.y {
		newPos.maxY = newPos.y
	} else {
		newPos.maxY = p.maxY
	}

	return newPos
}

func InTargetArea(p Position, t Target) bool {
	if t.x0 <= p.x && p.x <= t.x1 {
		if t.y0 >= p.y && p.y >= t.y1 {
			return true
		}
	}

	return false
}

func PassedTargetArea(p Position, v Velocity, t Target) bool {
	if p.x > t.x1 || p.y < t.y1 {
		return true
	}
	return false
}

func StuckX(p Position, v Velocity, t Target) bool {
	if v.x == 0 {
		if !(t.x0 <= p.x && p.x <= t.x1) {
			return true
		}
	}

	return false
}

func WillHit(pos Position, v Velocity, target Target) (bool, Position) {
	hit := false
	for {
		if StuckX(pos, v, target) {
			break
		}
		if PassedTargetArea(pos, v, target) {
			break
		}

		if InTargetArea(pos, target) {
			hit = true
			break
		}

		pos = ApplyVelocity(v, pos)
		v = UpdateVelocity(v)
	}
	return hit, pos
}

func FindMaxY(target Target) (int, int) {
	bestPos := Position{}
	hitCount := 0

	for i := 1; i < 400; i++ {
		for j := -300; j < 400; j++ {
			pos := Position{}
			vel := Velocity{i, j}
			hit, endPos := WillHit(pos, vel, target)
			if hit {
				hitCount += 1
			}
			if hit && endPos.maxY > bestPos.maxY {
				bestPos = endPos
			}
		}
	}

	return bestPos.maxY, hitCount
}
