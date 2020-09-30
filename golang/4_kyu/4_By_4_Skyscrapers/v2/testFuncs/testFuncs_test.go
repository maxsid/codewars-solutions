package testFuncs

import (
	"reflect"
	"testing"
)

func Test_setSkyscrapers(t *testing.T) {
	var (
		right = [2]int{0, 1}
		down = [2]int{1, 0}
		left = [2]int{0, -1}
		up = [2]int{-1, 0}
	)
	type args struct {
		town        [][]int
		dir         [2]int
		pos         [2]int
		skyscrapers []int
	}
	tests := []struct {
		name string
		args args
		wantS [][]int
	}{
		{
			name: "3 Down from 0;0",
			args: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			}, down, [2]int{0, 0}, []int{2, 3, 4}},
			wantS: [][]int{
				{2, 0, 0, 0},
				{3, 0, 0, 0},
				{4, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "4 Left from 2;3",
			args: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			}, left, [2]int{2, 3}, []int{1, 2, 3, 4}},
			wantS: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{4, 3, 2, 1},
				{0, 0, 0, 0},
			},
		},
		{
			name: "1 Up from 3;1",
			args: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			}, up, [2]int{3, 1}, []int{4}},
			wantS: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 4, 0, 0},
			},
		},
		{
			name: "2 Right from 3;0",
			args: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			}, right, [2]int{3, 0}, []int{3, 4}},
			wantS: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{3, 4, 0, 0},
			},
		},
		{
			name: "4 Right from 3;2",
			args: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			}, right, [2]int{3, 2}, []int{1, 2, 3, 4}},
			wantS: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 2},
			},
		},
		{
			name: "Mustn't change if it's not empty. 4 Right from 3;0",
			args: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 4, 0},
			}, right, [2]int{3, 0}, []int{1, 2, 3, 4}},
			wantS: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 2, 4, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setSkyscrapers(tt.args.town, tt.args.dir, tt.args.pos, tt.args.skyscrapers)
			if !reflect.DeepEqual(tt.args.town, tt.wantS) {
				t.Errorf("setSkyscrapers() = %v, want %v", tt.args.town, tt.wantS)
			}
		})
	}
}


func Test_intersection(t *testing.T) {
	type args struct {
		a        [][]int
		b        [][]int
		reverseB bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "None",
			args: args{
				a:        [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				b:        [][]int{
					{1, 3, 2, 4},
					{2, 4, 3, 1},
					{3, 1, 4, 2},
					{1, 4, 2, 3},
				},
				reverseB: false,
			},
			want: [][]int{},
		},
		{
			name: "None reversed",
			args: args{
				a:        [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				b:        [][]int{
					{1, 3, 2, 4},
					{2, 4, 3, 1},
					{3, 4, 1, 2},
					{1, 4, 2, 3},
				},
				reverseB: true,
			},
			want: [][]int{},
		},
		{
			name: "Simple two findings",
			args: args{
				a:        [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				b:        [][]int{
					{1, 3, 2, 4},
					{2, 4, 3, 1},
					{2, 3, 4, 1},
					{1, 2, 3, 4},
				},
				reverseB: false,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 1},
			},
		},
		{
			name: "Simple two reversedB findings",
			args: args{
				a:        [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				b:        [][]int{
					{1, 3, 2, 4},
					{2, 1, 4, 3},
					{2, 4, 3, 1},
					{4, 3, 2, 1},
				},
				reverseB: true,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{3, 4, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.a, tt.args.b, tt.args.reverseB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterByMustContain(t *testing.T) {
	type args struct {
		variants    [][]int
		mustContain [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "None",
			args: args{
				variants:    [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				mustContain: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{2, 3},
				},
			},
			want: [][]int{},
		},
		{
			name: "Simple",
			args: args{
				variants:    [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				mustContain: [][]int{
					{1, 2, 3},
					nil,
					{3, 1},
					nil,
				},
			},
			want: [][]int{
				{1, 2, 3, 4},
				{3, 4, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterByMustContain(tt.args.variants, tt.args.mustContain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterByMustContain() = %v, want %v", got, tt.want)
			}
		})
	}
}