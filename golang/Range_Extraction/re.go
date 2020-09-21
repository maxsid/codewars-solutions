package kata

import (
	"fmt"
	"strings"
)

func checkRange(inRange []int, inc *int, end bool) (out string, outRange []int) {
	outRange = inRange
	l := len(outRange)
	if !end && *inc >= 0 && outRange[l-1] == outRange[l-2]+1 {
		*inc = 1
	} else if !end && *inc <= 0 && outRange[l-1] == outRange[l-2]-1 {
		*inc = -1
	} else if l == 2 {
		out = fmt.Sprint(outRange[0])
		outRange = outRange[1:]
	} else if l == 3 {
		out = fmt.Sprintf("%d,%d", outRange[0], outRange[1])
		outRange = outRange[2:]
	} else {
		out = fmt.Sprintf("%d-%d", outRange[0], outRange[l-2])
		outRange = []int{outRange[l-1]}
	}
	return
}

func Solution(list []int) string {
	result := make([]string, 0)
	lRange := []int{list[0]}
	var out string
	inc := 0
	for _, v := range list[1:] {
		lRange = append(lRange, v)
		if out, lRange = checkRange(lRange, &inc, false); out != "" {
			result = append(result, out)
		}
	}
	lRange = append(lRange, 0)
	out, _ = checkRange(lRange, &inc, true)
	result = append(result, out)
	return strings.Join(result, ",")
}
