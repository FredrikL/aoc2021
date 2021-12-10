package aoc

func flashDumbo(flashState [][]bool, state [][]int, i, j int) [][]int {
	if state[i][j] < 10 || flashState[i][j] {
		return state
	}
	flashState[i][j] = true
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x < 0 {
				continue
			}
			if x >= len(state) {
				continue
			}
			if y >= len(state[i]) {
				continue
			}
			if y < 0 {
				continue
			}
			if x == i && y == j {
				continue
			}

			state[x][y] += 1

			flashDumbo(flashState, state, x, y)

		}
	}

	return state
}

func StepDumbo(state [][]int) ([][]int, int) {
	flashState := reset(state)
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			state[i][j] += 1
		}
	}

	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			state = flashDumbo(flashState, state, i, j)
		}
	}

	flashSum := 0
	// reset
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			if state[i][j] > 9 {
				state[i][j] = 0
				flashSum++
			}
		}
	}

	return state, flashSum
}

func reset(state [][]int) [][]bool {
	flashState := make([][]bool, len(state))
	for i := 0; i < len(state); i++ {
		flashState[i] = make([]bool, len(state))
	}
	return flashState
}
