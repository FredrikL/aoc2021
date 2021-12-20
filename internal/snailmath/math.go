package snailmath

import "strconv"

type Pair struct {
	left, right *Pair
	v           int
}

func ParsePairs(s string, pos int) (Pair, int) {
	if s[pos] == '[' {
		l, npos := ParsePairs(s, pos+1)
		r, npos := ParsePairs(s, npos+1)
		return Pair{
			left:  &l,
			right: &r,
		}, npos + 1
	}

	v, _ := strconv.Atoi(s[pos : pos+1])
	return Pair{
		v: v,
	}, pos + 1
}

func AddPairs(p1, p2 Pair) Pair {
	return Pair{
		left:  &p1,
		right: &p2,
	}
}

func ReducePair(pair Pair) Pair {
	return Pair{}
}

func explode(pair *Pair, depth int) bool {
	var l, r bool
	if depth == 3 {

		return true
	}

	if pair.left != nil {
		l = explode(pair.left, depth+1)
	}
	if pair.right != nil {
		r = explode(pair.right, depth+1)
	}

	return l || r
}

func split(pair Pair) bool {
	return false
}
