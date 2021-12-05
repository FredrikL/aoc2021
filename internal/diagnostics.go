package aoc

import "strconv"

func GetMostCommonForColumn(lines []string, index int) string {
	z, o := 0, 0

	for i := 0; i < len(lines); i++ {
		if lines[i][index] == '1' {
			o++
		} else {
			z++
		}
	}

	if o >= z {
		return "1"
	}
	return "0"
}

func GetLeastCommonForColumn(lines []string, index int) string {
	z, o := 0, 0

	for i := 0; i < len(lines); i++ {
		if lines[i][index] == '1' {
			o++
		} else {
			z++
		}
	}

	if z <= o {
		return "0"
	}
	return "1"
}

func GetGammaRate(lines []string) int64 {
	cols := len(lines[0])
	result := ""

	for i := 0; i < cols; i++ {
		v := GetMostCommonForColumn(lines, i)
		result = result + v
	}

	r, _ := strconv.ParseInt(result, 2, 64)

	return r
}

func GetEpsilonRate(lines []string) int64 {
	cols := len(lines[0])
	result := ""

	for i := 0; i < cols; i++ {
		v := GetLeastCommonForColumn(lines, i)
		result = result + v
	}

	r, _ := strconv.ParseInt(result, 2, 64)

	return r
}

func GetOxygenRating(lines []string) int64 {
	cols := len(lines[0])
	for i := 0; i < cols; i++ {
		v := GetMostCommonForColumn(lines, i)
		var l []string
		for j := 0; j < len(lines); j++ {
			if string(lines[j][i]) == v {
				l = append(l, lines[j])
			}
		}
		lines = l
	}

	r, _ := strconv.ParseInt(lines[0], 2, 64)
	return r
}

func GetCo2Rating(lines []string) int64 {
	cols := len(lines[0])
	for i := 0; i < cols; i++ {
		v := GetLeastCommonForColumn(lines, i)
		var l []string
		for j := 0; j < len(lines); j++ {
			if string(lines[j][i]) == v {
				l = append(l, lines[j])
			}
		}
		lines = l
		if len(lines) == 1 {
			r, _ := strconv.ParseInt(lines[0], 2, 64)
			return r
		}

	}

	return 0
}
