package kata

const size = 7
var allFloorNums = []int{1, 2, 3, 4, 5, 6, 7}

func SolvePuzzle(clues []int) [][]int {
	all := permutations(allFloorNums, 0)
	allLines := make([][][]int, len(clues)/2)
	for i := 0; i < len(clues)/2; i++ {
		clueS, clueE := getClues(clues, i)
		vars := filterVarsByClue(clueS, all, false)
		vars = filterVarsByClue(clueE, vars, true)
		allLines[i] = vars
	}
	vertLines, horizLines := allLines[:size], allLines[size:]
	for !allHaveOneLen(horizLines) {
		for i := range horizLines {
			horizLines[i] = pickUpLine(vertLines, horizLines, i)
		}
		for i := range vertLines {
			vertLines[i] = pickUpLine(horizLines, vertLines, i)
		}
	}
	out := make([][]int, size)
	for i, hl := range horizLines {
		out[i] = hl[0]
	}
	return out
}

func allHaveOneLen(x [][][]int) bool {
	for _, xl := range x {
		if len(xl) > 1 {
			return false
		}
	}
	return true
}

func pickUpLine(vertLines, horizLines [][][]int, n int) [][]int {
	possibleNums := make([][]int, size)
	for i := 0; i < size; i++ {
		canBeMap := make(map[int]struct{})
		for _, vl := range vertLines[i] {
			canBeMap[vl[n]] = struct{}{}
		}
		possibleNums[i] = make([]int, 0)
		for k := range canBeMap {
			possibleNums[i] = append(possibleNums[i], k)
		}
	}
	return mustContain(horizLines[n], possibleNums, false)
}

func permutations(a []int, i int) (out [][]int) {
	if i > len(a) {
		res := make([]int, len(a))
		copy(res, a)
		return [][]int{res}
	}
	out = permutations(a, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		out = append(out, permutations(a, i+1)...)
		a[i], a[j] = a[j], a[i]
	}
	return
}

func getClues(clues []int, n int) (s, e int) {
	s = clues[n]
	e = clues[size*(3+n/size)-(n%size+1)]
	if n >= size {
		s, e = e, s
	}
	return
}

func filterVarsByClue(clue int, variants [][]int, isEndClue bool) [][]int {
	if clue == 0 {
		return variants
	}
	out := make([][]int, 0)
	for _, p := range variants {
		v := p
		if isEndClue {
			v = reverseIntSlice(p)
		}
		see, biggest := 1, v[0]
		for i := 1; i < len(v); i++ {
			s := v[i]
			if s == size {
				see++
				break
			} else if s > biggest {
				biggest = s
				see++
			}
		}
		if see == clue {
			out = append(out, p)
		}
	}
	return out
}

func mustContain(variants [][]int, possibleNums [][]int, mustNot bool) [][]int {
	out := make([][]int, 0)
	nilSkips := 0
	for _, variant := range variants {
		isCorrect := true
		for i := 0; i < len(variant) && isCorrect; i++ {
			if possibleNums[i] == nil || len(possibleNums[i]) == 0 {
				nilSkips++
				continue
			}
			isCorrect = isCorrect && (mustNot != bHasA(variant[i], possibleNums[i]))
		}
		if nilSkips == len(possibleNums) {
			return variants
		} else {
			nilSkips = 0
		}
		if isCorrect {
			out = append(out, variant)
		}
	}
	return out
}

func bHasA(a int, b []int) bool {
	for _, bv := range b {
		if a == bv {
			return true
		}
	}
	return false
}

func reverseIntSlice(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}