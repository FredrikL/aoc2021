package aoc

type Dice struct {
	value  int
	thrown int
}

type Player struct {
	id       int
	position int
	score    int
}

func TakeTurn(dice *Dice, player *Player) {
	score := ThrowDice(dice) + ThrowDice(dice) + ThrowDice(dice)

	newPos := (player.position + score) % 10
	if newPos == 0 {
		newPos = 10
	}

	player.score += newPos
	player.position = newPos
}

func ThrowDice(dice *Dice) int {
	dice.thrown++
	dice.value++
	if dice.value > 100 {
		dice.value = 1
	}
	return dice.value
}
