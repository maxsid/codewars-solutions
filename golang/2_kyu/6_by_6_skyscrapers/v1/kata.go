package v1

const size = 6

func SolvePuzzle(clues []int) [][]int {
	town := initTown()
	all := permutations(getPossibleFloors(), 0)
	for i := 0; !isFilled(town); i = (i+1)%(size *2) {
		posses := getPosses(i)
		clueS, clueE := getClues(clues, i)
		vars := filterVarsByClue(clueS, all, false)
		vars = filterVarsByClue(clueE, vars, true)
		hasNums := getSkipNumsByPosses(town, posses)
		vars = filterByMustNotContain(vars, hasNums)
		if len(vars) == 0 {
			continue
		} else if len(vars) == 1 {
			setOnTown(town, posses, vars[0])
		} else {
			comm := commonNums(vars)
			setOnTown(town, posses, comm)
		}
	}
	return town
}

func initTown() [][]int {
	town := make([][]int, size)
	for i := 0; i < size; i++ {
		town[i] = make([]int, size)
	}
	return town
}

func getPossibleFloors() []int {
	possibleFloors := make([]int, size)
	for i := 0; i < size; i++ {
		possibleFloors[i] = i+1
	}
	return possibleFloors
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

// isFilled checks that the town has zeros
func isFilled(town [][]int) bool {
	for _, l := range town {
		for _, v := range l {
			if v == 0 {
				return false
			}
		}
	}
	return true
}

// getPosses returns all positions of a line in the town (if n < size it returns a vertical line, else horizontal one).
func getPosses(n int) [][2]int {
	out := make([][2]int, size)
	dir, pos := [2]int{1, 0}, [2]int{0, n}
	if n/size == 1 {
		dir, pos = [2]int{0, 1}, [2]int{n- size, 0}
	}
	for i := 0; i < size; i++ {
		out[i] = pos
		pos[0], pos[1] = pos[0]+dir[0], pos[1]+dir[1]
	}
	return out
}

// getClues returns clues of a line(s - from begin of the line, e - from end)
func getClues(clues []int, n int) (s, e int) {
	s = clues[n]
	e = clues[size*(3+n/size)-(n%size+1)]
	if n >= size {
		s, e = e, s
	}
	return
}

// filterVarsByClue returns variants of a line which corresponds to the clue()
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

// getExistsNums returns numbers which cannot be at the town[x][y] position
func getExistsNums(town [][]int, x, y int) []int {
	outMap := make(map[int]struct{})
	for i := 0; i < size; i++ {
		if town[x][i] != 0 {
			outMap[town[x][i]] = struct{}{}
		}
		if town[i][y] != 0 {
			outMap[town[i][y]] = struct{}{}
		}
	}
	out := make([]int, 0)
	for k := range outMap {
		out = append(out, k)
	}
	return out
}

func filterByMustNotContain(variants [][]int, mustNotContain [][]int) [][]int {
	out := make([][]int, 0)
	nilSkips := 0
	for _, variant := range variants {
		isCorrect := true
		for i := 0; i < len(variant) && isCorrect; i++ {
			if mustNotContain[i] == nil || len(mustNotContain[i]) == 0 {
				nilSkips++
				continue
			}
			isCorrect = isCorrect && !bHasA(variant[i], mustNotContain[i])
		}
		if nilSkips == len(mustNotContain) {
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

func getSkipNumsByPosses(town [][]int, posses [][2]int) [][]int {
	hasNums := make([][]int, len(posses))
	for i, p := range posses {
		if town[p[0]][p[1]] != 0 {
			hasNums[i] = nil
		} else {
			hasNums[i] = getExistsNums(town, p[0], p[1])
		}
	}
	return hasNums
}

func commonNums(x [][]int) []int {
	out := make([]int, len(x[0]))
	copy(out, x[0])
	zeros := 0
	for _, l := range x {
		for i, v := range l {
			if out[i] != 0 && out[i] != v {
				out[i] = 0
				zeros += 1
			}
			if zeros >= len(out) {
				return []int{}
			}
		}
	}
	return out
}

func setOnTown(town [][]int, posses [][2]int, sks []int) {
	for i, p := range posses {
		if i >= len(sks) {
			break
		} else if town[p[0]][p[1]] != 0 || sks[i] == 0 {
			continue
		}
		town[p[0]][p[1]] = sks[i]
	}
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