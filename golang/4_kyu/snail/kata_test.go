package kata

import (
	"reflect"
	"testing"
)

func TestSnail(t *testing.T) {
	type args struct {
		snaipMap [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Kata 1",
			args: args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Empty",
			args: args{[][]int{{}, {}, {}}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snail(tt.args.snaipMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Snail() = %v, want %v", got, tt.want)
			}
		})
	}
}
