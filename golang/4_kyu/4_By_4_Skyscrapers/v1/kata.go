package v1

const size int = 4
const (
	right int = iota
	down
	left
	up
)

var directions = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

func SolvePuzzle(clues []int) [][]int {
	town := initTown()
	// setClueSkyscrapers(town, clues)
	return town
}

func initTown() [][]int {
	town := make([][]int, size)
	for i := 0; i < size; i++ {
		town[i] = make([]int, size)
	}
	return town
}

func setClueSkyscrapers(town [][]int, clues []int) {
	pos := [2]int{0, 0}
	for i, d := range [][]int{[]int{right, down}, []int{down, left}, []int{left, up}, []int{up, right}} {
		for j := 0; j < size; j++ {
			ci := i*size + j
			moveD, setD := directions[d[0]], d[1]
			if clues[ci] > 0 {
				setSkyscrapers(town, setD, pos, getSkyscrapersByClue(clues[ci]))
			}
			nPos := [2]int{pos[0] + moveD[0], pos[1] + moveD[1]}
			if nPos[0] < size && nPos[0] >= 0 && nPos[1] < size && nPos[1] >= 0 {
				pos = nPos
			}
		}
	}
}

func setSkyscrapers(town [][]int, dir int, pos [2]int, skyscrapers []int) {
	d := directions[dir]
	for _, s := range skyscrapers {
		if len(town) <= pos[0] || len(town[pos[0]]) <= pos[1] || town[pos[0]][pos[1]] != 0 {
			break
		}
		town[pos[0]][pos[1]] = s
		pos[0], pos[1] = pos[0]+d[0], pos[1]+d[1]
	}
}

func getSkyscrapersByClue(clue int) (s []int) {
	if clue == 1 {
		return []int{size}
	} else if clue == size {
		s = make([]int, size)
		for i := 0; i < size; i++ {
			s[i] = i + 1
		}
		return
	}
	s = make([]int, 0)
	for i := size - clue + 1; i <= size; i++ {
		s = append(s, i)
	}
	return s
}

func possibleLinesByClue(clue int) [][]int {
	simpleLine := make([]int, size)
	for i := 0; i < size; i++ {
		simpleLine[i] = i + 1
	}
	perms := permutations(simpleLine, -1)
	out := make([][]int, 0)
	for _, p := range perms {
		see, biggest := 1, p[0]
		for _, s := range p[1:] {
			if s == 4 {
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

func permutations(nums []int, k int) [][]int {
	if k == -1 {
		k = len(nums)
	}
	nNums := make([]int, len(nums))
	copy(nNums, nums)
	if k == 1 {
		return [][]int{nNums}
	}

	out := permutations(nNums, k-1)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			nNums[i], nNums[k-1] = nNums[k-1], nNums[i]
		} else {
			nNums[0], nNums[k-1] = nNums[k-1], nNums[0]
		}
		out = append(out, permutations(nNums, k-1)...)
	}
	return out
}
