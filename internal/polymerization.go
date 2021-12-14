package aoc

import "strings"

func ParsePolyInput(input string) (string, map[string]string) {
	parts := strings.Split(input, "\n")
	res := make(map[string]string)
	for i := 2; i < len(parts); i++ {
		poly := strings.Split(parts[i], " -> ")
		res[poly[0]] = poly[1]
	}

	return parts[0], res
}

func StepPoly(current string, lookup map[string]string) string {
	res := ""
	for i := 0; i < (len(current) - 1); i++ {
		k := current[i:(i + 2)]
		v := lookup[k]
		res += current[i:(i+1)] + v
	}
	res += current[len(current)-1:]
	return res
}

func CountPloy(poly string) map[string]int {
	res := make(map[string]int)

	for i := 0; i < (len(poly)); i++ {
		k := poly[i:(i + 1)]
		res[k]++
	}

	return res
}
