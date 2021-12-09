package maps

type LowPoint struct {
	value int
	x     int
	y     int
}

func ReturnLowPoints(input [][]int) []LowPoint {
	ret := []LowPoint{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			v := input[i][j]

			if i > 0 && input[i-1][j] <= v {
				continue
			}
			if j > 0 && input[i][j-1] <= v {
				continue
			}

			if i < len(input)-1 && input[i+1][j] <= v {
				continue
			}

			if j < len(input[i])-1 && input[i][j+1] <= v {
				continue
			}

			p := LowPoint{
				value: v,
				x:     i,
				y:     j,
			}

			ret = append(ret, p)

		}

	}
	return ret
}

func SumResult(v []int) int {
	res := 0
	for _, v := range v {
		res += (v + 1)
	}
	return res
}

func FindBasinSize(input [][]int, x, y int) int {
	value := 0
	if input[x][y] != 9 {
		input[x][y] = 9
		value = 1
		if x > 0 {
			value += FindBasinSize(input, x-1, y)
		}
		if y > 0 {
			value += FindBasinSize(input, x, y-1)
		}
		if x < (len(input) - 1) {
			value += FindBasinSize(input, x+1, y)
		}
		if y < (len(input[x]) - 1) {
			value += FindBasinSize(input, x, y+1)
		}
	}
	return value
}
