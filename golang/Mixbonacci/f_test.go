package main

import (
	"reflect"
	"testing"
)

func Test_sequence_Next(t *testing.T) {
	type fields struct {
		l   []int64
		n   int
		nnf func([]int64) int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []int64
	}{
		{
			name: "Fibonacci numbers test",
			fields: fields{
				l:   []int64{0, 1},
				nnf: nextFibonacciNum,
			},
			want: []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155},
		},
		{
			name: "Padovan numbers test",
			fields: fields{
				l:   []int64{1, 0, 0},
				nnf: nextPadovanNum,
			},
			want: []int64{1, 0, 0, 1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 7, 9, 12, 16, 21, 28, 37, 49, 65, 86, 114, 151, 200, 265, 351, 465, 616, 816, 1081, 1432, 1897, 2513, 3329, 4410, 5842, 7739, 10252, 13581, 17991, 23833, 31572, 41824, 55405, 73396, 97229, 128801, 170625},
		},
		{
			name: "Jacobsthal numbers test",
			fields: fields{
				l:   []int64{0, 1},
				nnf: nextJacobsthalNum,
			},
			want: []int64{0, 1, 1, 3, 5, 11, 21, 43, 85, 171, 341, 683, 1365, 2731, 5461, 10923, 21845, 43691, 87381, 174763, 349525, 699051, 1398101, 2796203, 5592405, 11184811, 22369621, 44739243, 89478485, 178956971, 357913941, 715827883, 1431655765, 2863311531, 5726623061},
		},
		{
			name: "Pell numbers test",
			fields: fields{
				l:   []int64{0, 1},
				nnf: nextPellNum,
			},
			want: []int64{0, 1, 2, 5, 12, 29, 70, 169, 408, 985, 2378, 5741, 13860, 33461, 80782, 195025, 470832, 1136689, 2744210, 6625109, 15994428, 38613965, 93222358, 225058681, 543339720, 1311738121, 3166815962, 7645370045, 18457556052, 44560482149, 107578520350, 259717522849},
		},
		{
			name: "Tribonacci numbers test",
			fields: fields{
				l:   []int64{0, 0, 1},
				nnf: nextTribonacciNum,
			},
			want: []int64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136, 5768, 10609, 19513, 35890, 66012, 121415, 223317, 410744, 755476, 1389537, 2555757, 4700770, 8646064, 15902591, 29249425, 53798080, 98950096, 181997601, 334745777, 615693474, 1132436852},
		},
		{
			name: "Tetranacci numbers test",
			fields: fields{
				l:   []int64{0, 0, 0, 1},
				nnf: nextTetranacciNum,
			},
			want: []int64{0, 0, 0, 1, 1, 2, 4, 8, 15, 29, 56, 108, 208, 401, 773, 1490, 2872, 5536, 10671, 20569, 39648, 76424, 147312, 283953, 547337, 1055026, 2033628, 3919944, 7555935, 14564533, 28074040, 54114452, 104308960, 201061985, 387559437, 747044834, 1439975216, 2775641472},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sequence{
				l:   tt.fields.l,
				n:   tt.fields.n,
				nnf: tt.fields.nnf,
			}
			for _, want := range tt.want {
				if got := s.Next(); got != want {
					t.Errorf("sequence.Next() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMixbonacci(t *testing.T) {
	type args struct {
		pattern []string
		length  int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Edge Cases 1",
			args: args{[]string{}, 10},
			want: []int64{},
		},
		{
			name: "Edge Cases 2",
			args: args{[]string{"fib"}, 0},
			want: []int64{},
		},
		{
			name: "Multi-element Patterns 1",
			args: args{[]string{"fib", "tet"}, 10},
			want: []int64{0, 0, 1, 0, 1, 0, 2, 1, 3, 1},
		},
		{
			name: "Multi-element Patterns 2",
			args: args{[]string{"jac", "jac", "pel"}, 10},
			want: []int64{0, 1, 0, 1, 3, 1, 5, 11, 2, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mixbonacci(tt.args.pattern, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mixbonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
