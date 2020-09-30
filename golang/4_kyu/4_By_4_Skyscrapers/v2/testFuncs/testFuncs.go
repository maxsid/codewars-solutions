package testFuncs

import (
	"reflect"
)

func setSkyscrapers(town [][]int, dir, pos [2]int, skyscrapers []int) {
	for _, s := range skyscrapers {
		if len(town) <= pos[0] || len(town[pos[0]]) <= pos[1] || town[pos[0]][pos[1]] != 0 {
			break
		}
		town[pos[0]][pos[1]] = s
		pos[0], pos[1] = pos[0]+dir[0], pos[1]+dir[1]
	}
}

func intersection(a, b [][]int, reverseB bool) [][]int {
	out := make([][]int, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			rb := b[j]
			if reverseB {
				rb = reverseIntSlice(rb)
			}
			if reflect.DeepEqual(a[i], rb) {
				out = append(out, a[i])
			}
		}
	}
	return out
}


func filterByMustContain(variants [][]int, mustContain [][]int) [][]int {
	out := make([][]int, 0)
	for _, variant := range variants {
		isCorrect := true
		for i := 0; i < len(variant) && isCorrect; i++ {
			if mustContain[i] == nil {
				continue
			}
			//isCorrect = isCorrect && bHasA(variant[i], mustContain[i])
			isCorrect = isCorrect && bHasA(variant[i], mustContain[i])
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