package maps

import "strings"

func ParseInput(input string) map[string][]string {
	result := make(map[string][]string)
	parts := strings.Split(input, "\n")
	for _, l := range parts {
		nodes := strings.Split(l, "-")
		val := result[nodes[0]]
		val = append(val, nodes[1])
		result[nodes[0]] = val

		val = result[nodes[1]]
		val = append(val, nodes[0])
		result[nodes[1]] = val
	}
	return result
}

func includes(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TraverseMap(path []string, current string, graph map[string][]string) int {
	if current == "end" {
		return 1
	}

	path = append(path, current)

	next := graph[current]

	count := 0

	for _, n := range next {
		if strings.ToUpper(n) == n {
			count += TraverseMap(path, n, graph)
		} else if !includes(n, path) {
			count += TraverseMap(path, n, graph)
		}
	}

	return count
}

func anyDoubles(list []string) bool {
	m := map[string]bool{}

	for _, e := range list {
		if strings.ToUpper(e) != e {
			_, ok := m[e]
			if ok {
				return true
			}
			m[e] = true
		}
	}
	return false
}

func TraverseMapP2(path []string, current string, graph map[string][]string) int {
	if current == "end" {
		return 1
	}

	path = append(path, current)

	next := graph[current]

	count := 0

	for _, n := range next {
		if n == "start" {
			continue
		}
		if strings.ToUpper(n) == n {
			count += TraverseMapP2(path, n, graph)
		} else if !includes(n, path) {
			count += TraverseMapP2(path, n, graph)
		} else {
			if !anyDoubles(path) {
				count += TraverseMapP2(path, n, graph)
			}
		}
	}

	return count
}
