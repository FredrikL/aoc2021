package aoc

func CreateLanterns(input []int) map[int]int {
	m := make(map[int]int)

	for _, v := range input {
		m[v]++
	}

	return m
}

func StepLanterns(input map[int]int) map[int]int {
	m := make(map[int]int)

	for k, v := range input {
		if k > 0 {
			m[k-1] = v
		}
	}

	if input[0] > 0 {
		m[6] += input[0]
		m[8] = input[0]
	}
	return m
}
