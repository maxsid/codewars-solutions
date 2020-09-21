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
			name: "Sample test case",
			args: args{[]int{
				0, 0, 1, 2,
				0, 2, 0, 0,
				0, 3, 0, 0,
				0, 1, 0, 0}},
			want: [][]int{
				{2, 1, 4, 3},
				{3, 4, 1, 2},
				{4, 2, 3, 1},
				{1, 3, 2, 4},
			},
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

func Test_getSkyscrapersByClue(t *testing.T) {
	type args struct {
		clue int
	}
	tests := []struct {
		name  string
		args  args
		wantS []int
	}{
		{
			name:  "4 visible",
			args:  args{4},
			wantS: []int{1, 2, 3, 4},
		},
		{
			name:  "3 visible",
			args:  args{3},
			wantS: []int{2, 3, 4},
		},
		{
			name:  "2 visible",
			args:  args{2},
			wantS: []int{3, 4},
		},
		{
			name:  "1 visible",
			args:  args{1},
			wantS: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := getSkyscrapersByClue(tt.args.clue); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("getSkyscrapersByClue() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_setSkyscrapers(t *testing.T) {
	type args struct {
		town        [][]int
		dir         int
		pos         [2]int
		skyscrapers []int
	}
	tests := []struct {
		name  string
		args  args
		wantS [][]int
	}{
		{
			name: "3 Down from 0;0",
			args: args{initTown(), down, [2]int{0, 0}, []int{2, 3, 4}},
			wantS: [][]int{
				[]int{2, 0, 0, 0},
				[]int{3, 0, 0, 0},
				[]int{4, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
		},
		{
			name: "4 Left from 2;3",
			args: args{initTown(), left, [2]int{2, 3}, []int{1, 2, 3, 4}},
			wantS: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{4, 3, 2, 1},
				[]int{0, 0, 0, 0},
			},
		},
		{
			name: "1 Up from 3;1",
			args: args{initTown(), up, [2]int{3, 1}, []int{4}},
			wantS: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 4, 0, 0},
			},
		},
		{
			name: "2 Right from 3;0",
			args: args{initTown(), right, [2]int{3, 0}, []int{3, 4}},
			wantS: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{3, 4, 0, 0},
			},
		},
		{
			name: "4 Right from 3;2",
			args: args{initTown(), right, [2]int{3, 2}, []int{1, 2, 3, 4}},
			wantS: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 1, 2},
			},
		},
		{
			name: "Mustn't change if it's not empty. 4 Right from 3;0",
			args: args{[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 4, 0},
			}, right, [2]int{3, 0}, []int{1, 2, 3, 4}},
			wantS: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{1, 2, 4, 0},
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

func Test_setClueSkyscrapers(t *testing.T) {
	type args struct {
		town [][]int
		clue []int
	}
	tests := []struct {
		name  string
		args  args
		wantS [][]int
	}{
		{
			name: "Simple Test 1",
			args: args{initTown(), []int{
				2, 1, 0, 4,
				0, 0, 0, 0,
				0, 2, 1, 0,
				0, 3, 1, 2}},
			wantS: [][]int{
				[]int{3, 4, 0, 1},
				[]int{4, 0, 0, 2},
				[]int{2, 3, 4, 3},
				[]int{0, 4, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setClueSkyscrapers(tt.args.town, tt.args.clue)
			if !reflect.DeepEqual(tt.args.town, tt.wantS) {
				t.Errorf("setClueSkyscrapers() = %v, want %v", tt.args.town, tt.wantS)
			}
		})
	}
}
