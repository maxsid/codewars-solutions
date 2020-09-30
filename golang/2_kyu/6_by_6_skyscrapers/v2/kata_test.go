package kata

import (
	"reflect"
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	type args struct {
		clues []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Simple 1",
			args: args{[]int{
				3, 2, 2, 3, 2, 1,
				1, 2, 3, 3, 2, 2,
				5, 1, 2, 2, 4, 3,
				3, 2, 1, 2, 2, 4}},
			want: [][]int{
				{ 2, 1, 4, 3, 5, 6 },
				{ 1, 6, 3, 2, 4, 5 },
				{ 4, 3, 6, 5, 1, 2 },
				{ 6, 5, 2, 1, 3, 4 },
				{ 5, 4, 1, 6, 2, 3 },
				{ 3, 2, 5, 4, 6, 1 }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePuzzle(tt.args.clues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolvePuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mustContain(t *testing.T) {
	type args struct {
		variants     [][]int
		possibleNums [][]int
		mustNot      bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "None (false)",
			args: args{
				variants:    [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				possibleNums: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{2, 3},
				},
				mustNot: false,
			},
			want: [][]int{},
		},
		{
			name: "Simple (false)",
			args: args{
				variants:    [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				possibleNums: [][]int{
					{1, 2, 3},
					nil,
					{3, 1},
					nil,
				},
				mustNot: false,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{3, 4, 1, 2},
			},
		},
		{
			name: "None (true)",
			args: args{
				variants:       [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				possibleNums: [][]int{
					{1, 2},
					nil,
					{3, 1, 4},
					{3},
				},
				mustNot: true,
			},
			want: [][]int{},
		},
		{
			name: "Simple (true)",
			args: args{
				variants:       [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				possibleNums: [][]int{
					{1, 2},
					nil,
					{3, 4},
					{3},
				},
				mustNot: true,
			},
			want: [][]int{
				{3, 4, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mustContain(tt.args.variants, tt.args.possibleNums, tt.args.mustNot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mustContain() = %v, want %v", got, tt.want)
			}
		})
	}
}