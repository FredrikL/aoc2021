package snailmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadTestData(t *testing.T) {
	p, _ := ParsePairs("[1,2]", 0)

	assert.Equal(t, 1, p.left.v)
	assert.Equal(t, 2, p.right.v)

	p, _ = ParsePairs("[[1,2],3]", 0)

	assert.Equal(t, 0, p.left.v)
	assert.NotNil(t, p.left)
	assert.Equal(t, 3, p.right.v)

	p, _ = ParsePairs("[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]", 0)
	assert.NotNil(t, p.left)
	assert.NotNil(t, p.right)
}

func Test_AddPairs(t *testing.T) {
	p1, _ := ParsePairs("[1,2]", 0)

	p2, _ := ParsePairs("[[3,4],5]", 0)

	p := AddPairs(p1, p2)

	assert.NotNil(t, p.left)
	assert.NotNil(t, p.right)
	assert.Equal(t, 5, p.right.right.v)
}

func test_Explode(t *testing.T) {
	p, _ := ParsePairs("[[[[[9,8],1],2],3],4]", 0)

	exp := explode(&p, 0)

	assert.True(t, exp)
	assert.Equal(t, 0, p.left.left.left.left.v)
	assert.Equal(t, 9, p.left.left.left.right.v)
}
