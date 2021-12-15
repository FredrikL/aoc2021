package aoc

import "strings"

func ParsePolyInput(input string) (string, map[string]string, map[string]int) {
	parts := strings.Split(input, "\n")
	res := make(map[string]string)
	counts := make(map[string]int)
	template := parts[0]
	for i := 2; i < len(parts); i++ {
		poly := strings.Split(parts[i], " -> ")
		res[poly[0]] = poly[1]
	}

	for i := 0; i < (len(template) - 1); i++ {
		k := template[i:(i + 2)]
		counts[k]++
	}

	return parts[0], res, counts
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
