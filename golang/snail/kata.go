package kata

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func Snail(snaipMap [][]int) []int {
	if len(snaipMap[0]) == 0 {
		return []int{}
	}
	n, pos, dir, out := len(snaipMap), [2]int{0, 0}, directions[0], make([]int, 0)
	for tc, ndn := 0, 0; tc < 4; {
		if pos[0] >= n || pos[1] >= n || pos[0] < 0 || pos[1] < 0 || snaipMap[pos[0]][pos[1]] == -1 {
			pos[0], pos[1] = pos[0]-dir[0], pos[1]-dir[1]
			ndn, tc = (ndn+1)%4, tc+1
			dir = directions[ndn]
			pos[0], pos[1] = pos[0]+dir[0], pos[1]+dir[1]
		} else {
			snaipMap[pos[0]][pos[1]], tc, out = -1, 0, append(out, snaipMap[pos[0]][pos[1]])
			pos[0], pos[1] = pos[0]+dir[0], pos[1]+dir[1]
		}
	}
	return out
}
