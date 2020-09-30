package v1

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
