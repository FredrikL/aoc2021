package maps

import "strings"

func ValidateChunks(line string) (string, []string) {
	parts := strings.Split(line, "")
	stack := []string{}
	for _, v := range parts {
		if v == "(" || v == "[" || v == "{" || v == "<" {
			stack = append(stack, v)
			continue
		}
		last := stack[len(stack)-1]
		if last == "(" && v != ")" {
			return v, stack
		}
		if last == "{" && v != "}" {
			return v, stack
		}
		if last == "<" && v != ">" {
			return v, stack
		}
		if last == "[" && v != "]" {
			return v, stack
		}

		stack = stack[:len(stack)-1]
	}
	return "", stack
}

func ReverseStack(stack []string) []string {
	n := []string{}
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == "{" {
			n = append(n, "}")
		}
		if stack[i] == "[" {
			n = append(n, "]")
		}
		if stack[i] == "(" {
			n = append(n, ")")
		}
		if stack[i] == "<" {
			n = append(n, ">")
		}
	}
	return n
}

func CalcStackScore(stack []string) int {
	sum := 0
	for _, v := range stack {
		sum *= 5
		if v == "}" {
			sum += 3
		}
		if v == "]" {
			sum += 2
		}
		if v == ")" {
			sum += 1
		}
		if v == ">" {
			sum += 4
		}
	}
	return sum
}
