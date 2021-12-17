package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseExampleInput(t *testing.T) {
	target := ParseTargetArea("target area: x=20..30, y=-10..-5")

	assert.Equal(t, 20, target.x0)
	assert.Equal(t, 30, target.x1)
	assert.Equal(t, -5, target.y0)
	assert.Equal(t, -10, target.y1)
}

func Test_UpdateVelocity(t *testing.T) {
	v := Velocity{7, 2}
	pos := Position{}

	pos = ApplyVelocity(v, pos)
	assert.Equal(t, 7, pos.x)
	assert.Equal(t, 2, pos.y)

	v = UpdateVelocity(v)
	assert.Equal(t, 6, v.x)
	assert.Equal(t, 1, v.y)

	pos = ApplyVelocity(v, pos)
	assert.Equal(t, 13, pos.x)
	assert.Equal(t, 3, pos.y)

	v = UpdateVelocity(v)
	assert.Equal(t, 5, v.x)
	assert.Equal(t, 0, v.y)

	pos = ApplyVelocity(v, pos)
	assert.Equal(t, 18, pos.x)
	assert.Equal(t, 3, pos.y)
}

func Test_ShootFirstExample(t *testing.T) {
	target := ParseTargetArea("target area: x=20..30, y=-10..-5")
	v := Velocity{7, 2}
	pos := Position{}
	hit, _ := WillHit(pos, v, target)
	assert.True(t, hit)

	v = Velocity{6, 3}
	pos = Position{}
	hit, _ = WillHit(pos, v, target)
	assert.True(t, hit)

	v = Velocity{9, 0}
	pos = Position{}
	hit, _ = WillHit(pos, v, target)
	assert.True(t, hit)

	v = Velocity{17, -4}
	pos = Position{}
	hit, _ = WillHit(pos, v, target)
	assert.False(t, hit)
}

func Test_FindMaxY(t *testing.T) {
	target := ParseTargetArea("target area: x=20..30, y=-10..-5")

	maxY, hitCount := FindMaxY(target)

	assert.Equal(t, 45, maxY)
	assert.Equal(t, 112, hitCount)
}

func Test_FindMaxYDay17P1(t *testing.T) {
	// Brute force best force! :D
	target := ParseTargetArea("target area: x=288..330, y=-96..-50")

	maxY, hitCount := FindMaxY(target)

	assert.Equal(t, 4560, maxY)
	assert.Equal(t, 112, hitCount)
}
