package maps

import "container/heap"

type Pos struct {
	X, Y int
	val  int
}

type hp []Pos

func WalkMap(m, risk *[][]int, seen *[][]bool) {
	max_x := len(*m)
	max_y := len((*m)[0])

	h := make(hp, 1, 100)
	h[0] = Pos{0, 0, 0}

	for {
		pos := heap.Pop(&h).(Pos)

		x := pos.X
		y := pos.Y

		if x == max_x-1 && y == max_y-1 {
			return
		}

		if (*seen)[x][y] {
			continue
		}
		(*seen)[x][y] = true

		next := []Point{{X: x - 1, Y: y}, {X: x + 1, Y: y}, {X: x, Y: y - 1}, {X: x, Y: y + 1}}

		currentRisk := pos.val
		for _, v := range next {
			if v.X >= 0 && v.X < max_x && v.Y >= 0 && v.Y < max_y {
				newrisk := currentRisk + (*m)[v.X][v.Y]
				if newrisk >= (*risk)[v.X][v.Y] {
					continue
				}
				(*risk)[v.X][v.Y] = currentRisk + (*m)[v.X][v.Y]

				heap.Push(&h, Pos{v.X, v.Y, newrisk})
			}
		}
	}
}

func (h hp) Len() int { return len(h) }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h hp) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(Pos))
}
func (h *hp) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
