package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldParseExampleInputToMap(t *testing.T) {
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	expected := map[string][]string(map[string][]string{
		"A": {"start", "c", "b", "end"},
		"b": {"start", "A", "d", "end"},
		"c": {"A"}, "d": {"b"},
		"end":   {"A", "b"},
		"start": {"A", "b"}})
	res := ParseInput(input)

	assert.Equal(t, expected, res)
}

func Test_TraversFirstExample(t *testing.T) {
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	chart := ParseInput(input)

	paths := TraverseMap([]string{}, "start", chart)

	assert.Equal(t, 10, paths)
}

func Test_TraversSecondExample(t *testing.T) {
	input := `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
	chart := ParseInput(input)

	paths := TraverseMap([]string{}, "start", chart)

	assert.Equal(t, 19, paths)
}

func Test_TraversThirdExample(t *testing.T) {
	input := `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
	chart := ParseInput(input)

	paths := TraverseMap([]string{}, "start", chart)

	assert.Equal(t, 226, paths)
}

func Test_TraversDay12P1(t *testing.T) {
	input := `re-js
qx-CG
start-js
start-bj
qx-ak
js-bj
ak-re
CG-ak
js-CG
bj-re
ak-lg
lg-CG
qx-re
WP-ak
WP-end
re-lg
end-ak
WP-re
bj-CG
qx-start
bj-WP
JG-lg
end-lg
lg-iw`
	chart := ParseInput(input)

	paths := TraverseMap([]string{}, "start", chart)

	assert.Equal(t, 3230, paths)
}

func Test_TraversFirstExampleP2(t *testing.T) {
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	chart := ParseInput(input)

	paths := TraverseMapP2([]string{}, "start", chart)

	assert.Equal(t, 36, paths)
}

func Test_TraversSecondExampleP2(t *testing.T) {
	input := `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
	chart := ParseInput(input)

	paths := TraverseMapP2([]string{}, "start", chart)

	assert.Equal(t, 103, paths)
}

func Test_TraversThirdExampleP2(t *testing.T) {
	input := `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
	chart := ParseInput(input)

	paths := TraverseMapP2([]string{}, "start", chart)

	assert.Equal(t, 3509, paths)
}

func Test_TraversDay12P2(t *testing.T) {
	input := `re-js
qx-CG
start-js
start-bj
qx-ak
js-bj
ak-re
CG-ak
js-CG
bj-re
ak-lg
lg-CG
qx-re
WP-ak
WP-end
re-lg
end-ak
WP-re
bj-CG
qx-start
bj-WP
JG-lg
end-lg
lg-iw`
	chart := ParseInput(input)

	paths := TraverseMapP2([]string{}, "start", chart)

	assert.Equal(t, 83475, paths)
}
