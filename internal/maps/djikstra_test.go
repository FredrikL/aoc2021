package maps

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleMap string = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

const exampleBig string = `11637517422274862853338597396444961841755517295286
13813736722492484783351359589446246169155735727126
21365113283247622439435873354154698446526571955763
36949315694715142671582625378269373648937148475914
74634171118574528222968563933317967414442817852555
13191281372421239248353234135946434524615754563572
13599124212461123532357223464346833457545794456865
31254216394236532741534764385264587549637569865174
12931385212314249632342535174345364628545647573965
23119445813422155692453326671356443778246755488935
22748628533385973964449618417555172952866628316397
24924847833513595894462461691557357271266846838237
32476224394358733541546984465265719557637682166874
47151426715826253782693736489371484759148259586125
85745282229685639333179674144428178525553928963666
24212392483532341359464345246157545635726865674683
24611235323572234643468334575457944568656815567976
42365327415347643852645875496375698651748671976285
23142496323425351743453646285456475739656758684176
34221556924533266713564437782467554889357866599146
33859739644496184175551729528666283163977739427418
35135958944624616915573572712668468382377957949348
43587335415469844652657195576376821668748793277985
58262537826937364893714847591482595861259361697236
96856393331796741444281785255539289636664139174777
35323413594643452461575456357268656746837976785794
35722346434683345754579445686568155679767926678187
53476438526458754963756986517486719762859782187396
34253517434536462854564757396567586841767869795287
45332667135644377824675548893578665991468977611257
44961841755517295286662831639777394274188841538529
46246169155735727126684683823779579493488168151459
54698446526571955763768216687487932779859814388196
69373648937148475914825958612593616972361472718347
17967414442817852555392896366641391747775241285888
46434524615754563572686567468379767857948187896815
46833457545794456865681556797679266781878137789298
64587549637569865174867197628597821873961893298417
45364628545647573965675868417678697952878971816398
56443778246755488935786659914689776112579188722368
55172952866628316397773942741888415385299952649631
57357271266846838237795794934881681514599279262561
65719557637682166874879327798598143881961925499217
71484759148259586125936169723614727183472583829458
28178525553928963666413917477752412858886352396999
57545635726865674683797678579481878968159298917926
57944568656815567976792667818781377892989248891319
75698651748671976285978218739618932984172914319528
56475739656758684176786979528789718163989182927419
67554889357866599146897761125791887223681299833479`

func toInputMap(input string) ([][]int, [][]int, [][]bool) {
	result := [][]int{}

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		r := []int{}
		rp := strings.Split(row, "")
		for _, rv := range rp {
			v, _ := strconv.Atoi(rv)
			r = append(r, v)
		}
		result = append(result, r)
	}

	risk := make([][]int, len(result))
	seen := make([][]bool, len(result))
	for i := range risk {
		risk[i] = make([]int, len(result[0]))
		for j := range risk[i] {
			risk[i][j] = 10000
		}
		seen[i] = make([]bool, len(result[0]))
	}

	return result, risk, seen
}

func load_day15() string {
	file, _ := os.Open("../../test/day15_input.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

func Test_LoadExamplData(t *testing.T) {
	mp, risk, _ := toInputMap(exampleMap)

	assert.Len(t, mp, 10)
	assert.Len(t, mp[1], 10)

	assert.Len(t, risk, 10)
	assert.Len(t, risk[1], 10)
	assert.Equal(t, 0, risk[1][0])
}

func Test_WalkExample(t *testing.T) {
	mp, risk, seen := toInputMap(exampleMap)
	mp[0][0] = 0
	WalkMap(&mp, &risk, &seen)

	assert.Equal(t, 40, risk[9][9])
}

func Test_WalkExampleBig(t *testing.T) {
	mp, risk, seen := toInputMap(exampleBig)
	mp[0][0] = 0
	WalkMap(&mp, &risk, &seen)

	assert.Equal(t, 315, risk[49][49])
}
func Test_WalkDay15P1(t *testing.T) {
	mp, risk, seen := toInputMap(load_day15())
	mp[0][0] = 0
	WalkMap(&mp, &risk, &seen)

	assert.Equal(t, 462, risk[len(mp)-1][len(mp[0])-1])
}

func toLargerMap(input string) ([][]int, [][]int, [][]bool) {
	base, _, _ := toInputMap(input)

	result := make([][]int, len(base)*5)
	for i := range result {
		result[i] = make([]int, len(base[0])*5)
	}

	x_base := len(base)
	y_base := len(base[0])

	for x, b := range base {
		for y := range b {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					r_x := x + x_base*i
					r_y := y + y_base*j

					result[r_x][r_y] = base[x][y] + i + j
					if result[r_x][r_y] > 9 {
						result[r_x][r_y] %= 9
					}
				}
			}
		}
	}

	risk := make([][]int, len(result))
	seen := make([][]bool, len(result))

	for i := range risk {
		risk[i] = make([]int, len(result[0]))
		for j := range risk[i] {
			risk[i][j] = 10000
		}
		seen[i] = make([]bool, len(result[0]))
	}

	return result, risk, seen
}

func Test_LoadExamplDatap2(t *testing.T) {
	mp, _, _ := toLargerMap(exampleMap)
	large, _, _ := toInputMap(exampleBig)

	assert.Len(t, mp, 50)
	assert.Len(t, mp[1], 50)
	assert.Equal(t, large, mp)
}

func Test_WalkDay15P2(t *testing.T) {
	mp, risk, seen := toLargerMap(load_day15())
	mp[0][0] = 0
	WalkMap(&mp, &risk, &seen)

	assert.Equal(t, 2846, risk[len(mp)-1][len(mp[0])-1])
}
