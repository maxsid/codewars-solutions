package kata

func Encode(s string,n int) string {
	if s == "" {
		return ""
	} else if n >= len(s) {
		return s
	}
	rails, out := getRailsFromDecodedStr(s, n), ""
	for _, r := range rails {
		out += string(r)
	}
	return out
}

func Decode(s string,n int) string {
	if s == "" {
		return ""
	} else if n >= len(s) {
		return s
	}
	rails, rcis, out := getRailsFromEncodedStr(s, n), make([]int, n), make([]byte, 0)
	for i, r, dir := 0, 0, 1; i < len(s); r, i = r+dir, i+1 {
		rcis[r], out = rcis[r]+1, append(out, rails[r][rcis[r]])
		if r+dir >= n {
			dir = -1
		} else if r+dir < 0 {
			dir = 1
		}
	}
	return string(out)
}

func getRailsFromDecodedStr(s string, n int) [][]byte {
	rails := make([][]byte, n)
	for si, ri, dir := 0, 0, 1; si < len(s); si, ri = si+1, ri+dir {
		rails[ri] = append(rails[ri], s[si])
		if ri+dir >= n {
			dir = -1
		} else if ri+dir < 0 {
			dir = 1
		}
	}
	return rails
}

func getRailsFromEncodedStr(s string, n int) [][]byte {
	si, rails := 0, getRailsFromDecodedStr(s, n)
	for i := 0; i < len(rails); i++ {
		for j := 0; j < len(rails[i]); j++ {
			rails[i][j] = s[si]
			si++
		}
	}
	return rails
}