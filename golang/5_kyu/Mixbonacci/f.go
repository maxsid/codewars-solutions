package main

func main() {
}

func Mixbonacci(pattern []string, length int) []int64 {
	if len(pattern) == 0 {
		return []int64{}
	}
	seqs := map[string]*sequence{
		"fib": newSequence([]int64{0, 1}, nextFibonacciNum),
		"pad": newSequence([]int64{1, 0, 0}, nextPadovanNum),
		"jac": newSequence([]int64{0, 1}, nextJacobsthalNum),
		"pel": newSequence([]int64{0, 1}, nextPellNum),
		"tri": newSequence([]int64{0, 0, 1}, nextTribonacciNum),
		"tet": newSequence([]int64{0, 0, 0, 1}, nextTetranacciNum),
	}
	res := make([]int64, length)
	for pi, li := 0, 0; li < length; pi, li = pi+1, li+1 {
		if pi >= len(pattern) {
			pi = 0
		}
		res[li] = seqs[pattern[pi]].Next()
	}
	return res
}

type sequence struct {
	l   []int64
	n   int
	nnf func([]int64) int64
}

func newSequence(firstNums []int64, nextNumFunc func([]int64) int64) *sequence {
	return &sequence{
		l:   firstNums,
		n:   0,
		nnf: nextNumFunc,
	}
}

func (s *sequence) Next() int64 {
	if s.n < len(s.l) {
		s.n++
		return s.l[s.n-1]
	}
	nn := s.nnf(s.l)
	for i := 0; i < len(s.l)-1; i++ {
		s.l[i] = s.l[i+1]
	}
	s.l[len(s.l)-1] = nn
	return nn
}

func nextFibonacciNum(l []int64) int64 {
	return l[0] + l[1]
}

func nextPadovanNum(l []int64) int64 {
	return l[1] + l[0]
}

func nextPellNum(l []int64) int64 {
	return 2*l[1] + l[0]
}

func nextJacobsthalNum(l []int64) int64 {
	return l[1] + 2*l[0]
}

func nextTribonacciNum(l []int64) int64 {
	return l[2] + l[1] + l[0]
}

func nextTetranacciNum(l []int64) int64 {
	return l[0] + l[1] + l[2] + l[3]
}
