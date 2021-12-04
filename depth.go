package aoc

func DepthIncreseCount(mesurements []int) int {
	sum := 0
	for i := 1; i < len(mesurements); i++ {
		if mesurements[i-1] < mesurements[i] {
			sum += 1
		}
	}
	return sum
}

func MergeMesurments(mesurements []int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(mesurements)-2; i++ {
		ret = append(ret, mesurements[i]+mesurements[i+1]+mesurements[i+2])
	}
	return ret
}
