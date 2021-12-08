package aoc

import (
	"sort"
	"strings"
)

func CountDigitsInRow(input string) int {
	data := strings.Split(input, " | ")
	digits := strings.Split(data[1], " ")
	count := 0
	for _, d := range digits {
		switch len(d) {
		case 2:
			count++
		case 3:
			count++
		case 4:
			count++
		case 7:
			count++

		}
	}
	return count
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func minus(a, b []string) []string {

	var diff []string
	for _, x := range a {
		found := false
		for _, y := range b {
			if x == y {
				found = true
			}
		}
		if !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func MapDigits(input string) map[string]int {
	parts := strings.Split(input, " | ")
	data := strings.Split(parts[0], " ")
	lookup := make([]string, 10)

	len6 := []string{}
	len5 := []string{}

	for _, v := range data {
		switch len(v) {
		case 2:
			lookup[1] = v
		case 3:
			lookup[7] = v
		case 4:
			lookup[4] = v
		case 5:
			len5 = append(len5, v)
		case 6:
			len6 = append(len6, v)
		case 7:
			lookup[8] = v
		}
	}

	one := strings.Split(lookup[1], "")
	four := strings.Split(lookup[4], "")
	ef := difference(four, one)
	for _, v := range len6 {
		val := strings.Split(v, "")
		diff := difference(val, ef)
		if len(diff) == 5 {
			// 0
			lookup[0] = v
		} else {
			diff = difference(val, one)
			if len(diff) == 4 {
				// 9
				lookup[9] = v
			} else {
				lookup[6] = v
			}
		}
	}

	seven := strings.Split(lookup[7], "")

	for _, v := range len5 {
		val := strings.Split(v, "")
		diff := difference(val, seven)
		if len(diff) == 2 {
			lookup[3] = v
		} else {
			diff = minus(val, four)
			if len(diff) == 2 {
				lookup[5] = v
			} else {
				lookup[2] = v
			}
		}
	}

	res := make(map[string]int)
	for i, v := range lookup {
		res[v] = i
	}
	return res
}

func sortString(v string) string {
	p := strings.Split(v, "")
	sort.Strings(p)
	s := strings.Join(p, "")
	return s
}

func SortedKeys(input map[string]int) map[string]int {
	r := make(map[string]int)
	for k, v := range input {
		r[sortString(k)] = v
	}
	return r
}

func SortedArray(input []string) []string {
	r := make([]string, len(input))
	for i, v := range input {
		r[i] = sortString(v)
	}
	return r
}

func MapRow(input string) int {
	mapped := MapDigits(input)
	data := strings.Split(input, " | ")
	digits := strings.Split(data[1], " ")

	sortedMap := SortedKeys(mapped)
	sortedDigits := SortedArray(digits)

	value := (sortedMap[sortedDigits[0]] * 1000) + (sortedMap[sortedDigits[1]] * 100) + (sortedMap[sortedDigits[2]] * 10) + sortedMap[sortedDigits[3]]

	return value
}
