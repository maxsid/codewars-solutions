package kata

var directions = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

func CreateSpiral(n int) [][]int {
	if n < 1 {
		return nil
	} else if n == 1 {
		return [][]int{{1}}
	}
	sq := make([][]int, n)
	for i := 0; i < n; i++ {
		sq[i] = make([]int, n)
	}
	isDoubleRepeat, dirN, pos := false, 0, [2]int{0, 0}
	sq[0][0] = 1
	for i := 2; ; i++ {
		dir := directions[dirN]
		nPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if nPos[0] >= n || nPos[1] >= n || nPos[0] < 0 || nPos[1] < 0 || sq[nPos[0]][nPos[1]] != 0 {
			if isDoubleRepeat {
				break
			}
			isDoubleRepeat = true
			dirN = (dirN + 1) % 4
			i--
			continue
		}
		isDoubleRepeat, pos = false, nPos
		sq[pos[0]][pos[1]] = i
	}
	return sq
}
