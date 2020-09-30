package kata

import (
	"reflect"
	"testing"
)

func TestCreateSpiral(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	name: "calculates n < 1",
		// 	args: args{0},
		// 	want: nil,
		// },
		// {
		// 	name: "calculates n = 1",
		// 	args: args{1},
		// 	want: [][]int{{1}},
		// },
		{
			name: "calculates n = 2",
			args: args{2},
			want: [][]int{{1, 2}, {4, 3}},
		},
		{
			name: "calculates n = 3",
			args: args{3},
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSpiral(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSpiral() = %v, want %v", got, tt.want)
			}
		})
	}
}
