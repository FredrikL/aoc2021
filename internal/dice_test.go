package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// example data
// Player 1 starting position: 4
// Player 2 starting position: 8

func Test_PlayExampleGame(t *testing.T) {
	dice := Dice{}

	players := [2]Player{
		Player{1, 4, 0}, Player{2, 8, 0},
	}

	TakeTurn(&dice, &players[0])

	assert.Equal(t, 3, dice.value)
	assert.Equal(t, 3, dice.thrown)
	assert.Equal(t, 10, players[0].score)

	TakeTurn(&dice, &players[1])
	assert.Equal(t, 6, dice.value)
	assert.Equal(t, 6, dice.thrown)
	assert.Equal(t, 3, players[1].score)

	TakeTurn(&dice, &players[0])
	assert.Equal(t, 14, players[0].score)

	TakeTurn(&dice, &players[1])
	assert.Equal(t, 9, players[1].score)

	i := 0
	for {
		TakeTurn(&dice, &players[i])
		if players[i].score >= 1000 {
			break
		}
		i++
		if i == 2 {
			i = 0
		}
	}

	assert.Equal(t, 1000, players[0].score)
	assert.Equal(t, 745, players[1].score)
	assert.Equal(t, 993, dice.thrown)
}

/*
Player 1 starting position: 4
Player 2 starting position: 6
*/

func Test_PlayDay21P1(t *testing.T) {
	dice := Dice{}

	players := [2]Player{
		Player{1, 4, 0}, Player{2, 6, 0},
	}
	i := 0
	for {
		TakeTurn(&dice, &players[i])
		if players[i].score >= 1000 {
			break
		}
		i++
		if i == 2 {
			i = 0
		}
	}

	assert.Equal(t, 1000, players[0].score)
	assert.Equal(t, 895, players[1].score)
	assert.Equal(t, 993, dice.thrown)
	assert.Equal(t, 888735, players[1].score*dice.thrown)
}
