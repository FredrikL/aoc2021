package control

func GetMinMax(input []int) (int, int) {
	min, max := 0, 0
	for idx, v := range input {
		if idx == 0 {
			min = v
			max = v
			continue
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func fib(v int) int {
	if v == 0 {
		return 0
	}
	return v + fib(v-1)
}

func CalcFuel(input []int, pos int) int {
	sum := int(0)
	for _, v := range input {
		use := Abs(v - pos)
		sum += fib(use)
	}
	return sum
}

func FindBestPos(input []int, min, max int) int {
	bestUse := int(0)

	for i := min; i <= max; i++ {
		use := CalcFuel(input, i)

		if use < bestUse || bestUse == 0 {
			bestUse = use
		}
	}

	return bestUse
}
