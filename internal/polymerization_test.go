package aoc

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInputD14 string = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

const day14Input string = `KKOSPHCNOCHHHSPOBKVF

NV -> S
OK -> K
SO -> N
FN -> F
NB -> K
BV -> K
PN -> V
KC -> C
HF -> N
CK -> S
VP -> H
SK -> C
NO -> F
PB -> O
PF -> P
VC -> C
OB -> S
VF -> F
BP -> P
HO -> O
FF -> S
NF -> B
KK -> C
OC -> P
OV -> B
NK -> B
KO -> C
OH -> F
CV -> F
CH -> K
SC -> O
BN -> B
HS -> O
VK -> V
PV -> S
BO -> F
OO -> S
KB -> N
NS -> S
BF -> N
SH -> F
SB -> S
PP -> F
KN -> H
BB -> C
SS -> V
HP -> O
PK -> P
HK -> O
FH -> O
BC -> N
FK -> K
HN -> P
CC -> V
FO -> F
FP -> C
VO -> N
SF -> B
HC -> O
NN -> K
FC -> C
CS -> O
FV -> P
HV -> V
PO -> H
BH -> F
OF -> P
PC -> V
CN -> O
HB -> N
CF -> P
HH -> K
VH -> H
OP -> F
BK -> S
SP -> V
BS -> V
VB -> C
NH -> H
SN -> K
KH -> F
OS -> N
NP -> P
VN -> V
KV -> F
KP -> B
VS -> F
NC -> F
ON -> S
FB -> C
SV -> O
PS -> K
KF -> H
CP -> H
FS -> V
VV -> H
CB -> P
PH -> N
CO -> N
KS -> K`

func Test_ShouldHanldeD14Example(t *testing.T) {
	start, mp, _ := ParsePolyInput(exampleInputD14)

	assert.Equal(t, "NNCB", start)
	// assert.Equal(t, "NNCB", mp)

	step := StepPoly(start, mp)
	assert.Equal(t, "NCNBCHB", step)

	step = StepPoly(step, mp)
	assert.Equal(t, "NBCCNBBBCBHCB", step)

	step = StepPoly(step, mp)
	assert.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", step)

	step = StepPoly(step, mp)
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", step)

	// validate length
	step = start
	for i := 0; i < 10; i++ {
		step = StepPoly(step, mp)
	}
	assert.Len(t, step, 3073)

	c := CountPloy(step)
	assert.Equal(t, 1749, c["B"])

	ints := []int{}
	for _, v := range c {
		ints = append(ints, v)
	}
	sort.Ints(ints)
	assert.Equal(t, 1588, ints[len(ints)-1]-ints[0])
}

func Test_ShouldHanldeD14P1(t *testing.T) {
	start, mp, _ := ParsePolyInput(day14Input)

	assert.Equal(t, "KKOSPHCNOCHHHSPOBKVF", start)

	// validate length
	step := start
	for i := 0; i < 10; i++ {
		step = StepPoly(step, mp)
	}

	c := CountPloy(step)

	ints := []int{}
	for _, v := range c {
		ints = append(ints, v)
	}
	sort.Ints(ints)
	assert.Equal(t, 2321, ints[len(ints)-1]-ints[0])
}

func Test_ShouldHanldeD14P1P2Way(t *testing.T) {
	template, mp, counts := ParsePolyInput(exampleInputD14)

	letterCount := make(map[string]int)

	assert.Equal(t, "NNCB", template)

	for i := 0; i < 10; i++ {
		newcounts := make(map[string]int)
		for key, c := range counts {
			n := mp[key]
			parts := strings.Split(key, "")
			newcounts[parts[0]+n] += c
			newcounts[n+parts[1]] += c
			letterCount[n] += c
		}
		counts = newcounts
	}

	p := strings.Split(template, "")
	for _, v := range p {
		letterCount[v]++
	}

	ints := []int{}
	for _, v := range letterCount {
		ints = append(ints, v)
	}
	sort.Ints(ints)
	assert.Equal(t, 1588, ints[len(ints)-1]-ints[0])

}

func Test_ShouldHanldeD14ExampleCombinations(t *testing.T) {
	template, mp, counts := ParsePolyInput(exampleInputD14)

	letterCount := make(map[string]int)

	assert.Equal(t, "NNCB", template)

	for i := 0; i < 40; i++ {
		newcounts := make(map[string]int)
		for key, c := range counts {
			n := mp[key]
			parts := strings.Split(key, "")
			newcounts[parts[0]+n] += c
			newcounts[n+parts[1]] += c
			letterCount[n] += c
		}
		counts = newcounts
	}

	p := strings.Split(template, "")
	for _, v := range p {
		letterCount[v]++
	}

	ints := []int{}
	for _, v := range letterCount {
		ints = append(ints, v)
	}
	sort.Ints(ints)
	assert.Equal(t, 2188189693529, ints[len(ints)-1]-ints[0])

}

func Test_ShouldHanldeD14P2(t *testing.T) {
	template, mp, counts := ParsePolyInput(day14Input)

	letterCount := make(map[string]int)

	for i := 0; i < 40; i++ {
		newcounts := make(map[string]int)
		for key, c := range counts {
			n := mp[key]
			parts := strings.Split(key, "")
			newcounts[parts[0]+n] += c
			newcounts[n+parts[1]] += c
			letterCount[n] += c
		}
		counts = newcounts
	}

	p := strings.Split(template, "")
	for _, v := range p {
		letterCount[v]++
	}

	ints := []int{}
	for _, v := range letterCount {
		ints = append(ints, v)
	}
	sort.Ints(ints)
	assert.Equal(t, 2399822193707, ints[len(ints)-1]-ints[0])

}
