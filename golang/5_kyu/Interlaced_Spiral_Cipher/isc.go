package isc

import (
	"math"
	"strings"
)

var directions = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

func getNextPosition(pos [2]int, dirN, squareSize, step int) (int, [2]int) {
	dir := directions[dirN]
	for i := step; i > 0; i-- {
		nPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if nPos[0] >= squareSize || nPos[1] >= squareSize || nPos[0] < 0 || nPos[1] < 0 {
			return getNextPosition(pos, (dirN+1)%4, squareSize, i)
		}
		pos = nPos
	}
	return dirN, pos
}

func getSquareSize(sLen int) int          { return int(math.Ceil(math.Sqrt(float64(sLen)))) }
func getCirclesNumber(squareSize int) int { return int(math.Ceil(float64(squareSize) / 2)) }

func iscWalk(squareSize, circlesNumber int, do func(sqPos [2]int, sInd int)) {
	si := 0
	for circle := 0; circle < circlesNumber; circle++ {
		circleSize := squareSize - circle*2
		circleStartPosition := [2]int{circle, circle}
		step := circleSize - 1
		dirN := 0
		for csp, firstRun := 0, true; csp < step || firstRun; csp, firstRun = csp+1, false {
			for p, ok := [2]int{0, csp}, true; ok; ok, si = p != [2]int{0, csp}, si+1 {
				pos := [2]int{circleStartPosition[0] + p[0], circleStartPosition[1] + p[1]}
				do(pos, si)
				dirN, p = getNextPosition(p, dirN, circleSize, step)
			}
		}
	}
}

func encode(s string) string {
	sqSize := getSquareSize(len(s))
	circNums := getCirclesNumber(sqSize)
	square := make([][]byte, sqSize)
	for i := 0; i < sqSize; i++ {
		square[i] = make([]byte, sqSize)
	}
	iscWalk(sqSize, circNums, func(pos [2]int, sInd int) {
		if sInd < len(s) {
			square[pos[0]][pos[1]] = s[sInd]
		} else {
			square[pos[0]][pos[1]] = ' '
		}
	})
	resLine := ""
	for _, row := range square {
		resLine += string(row)
	}
	return resLine
}

func decode(s string) string {
	sqSize := getSquareSize(len(s))
	circNums := getCirclesNumber(sqSize)
	square := make([][]byte, sqSize)

	for i := 0; i < sqSize; i++ {
		square[i] = []byte(s[i*sqSize : (i+1)*sqSize])
	}
	resLine := ""
	iscWalk(sqSize, circNums, func(pos [2]int, _ int) {
		resLine += string(square[pos[0]][pos[1]])
	})

	return strings.TrimSpace(resLine)
}
