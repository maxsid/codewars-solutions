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
			name: "Kata",
			args: args{[]int{0,0,1,2,0,2,0,0,0,3,0,0,0,1,0,0}},
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


func Test_commonNums(t *testing.T) {
	type args struct {
		x [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Not found common numbers",
			args: args{x: [][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 1},
				{3, 4, 1, 2},
				{4, 1, 2, 3},
			}},
			want: []int{},
		},
		{
			name: "Simple found",
			args: args{x: [][]int{
				{1, 2, 3, 4},
				{2, 2, 4, 1},
				{3, 2, 1, 2},
				{4, 2, 2, 3},
			}},
			want: []int{0, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonNums(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseIntSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple Test",
			args: args{[]int{1, 2, 3, 4, 5, 6}},
			want: []int{6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseIntSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterByMustNotContain(t *testing.T) {
	type args struct {
		variants       [][]int
		mustNotContain [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "None",
			args: args{
				variants:       [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				mustNotContain: [][]int{
					{1, 2},
					nil,
					{3, 1, 4},
					{3},
				},
			},
			want: [][]int{},
		},
		{
			name: "Simple",
			args: args{
				variants:       [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				mustNotContain: [][]int{
					{1, 2},
					nil,
					{3, 4},
					{3},
				},
			},
			want: [][]int{
				{3, 4, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterByMustNotContain(tt.args.variants, tt.args.mustNotContain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterByMustNotContain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bHasA(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple True",
			args: args{
				a: 1,
				b: []int{1, 2, 3, 4},
			},
			want: true,
		},
		{
			name: "Simple False",
			args: args{
				a: 5,
				b: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bHasA(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("bHasA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleLinesByClue(t *testing.T) {
	type args struct {
		clue     int
		variants [][]int
		reverseVar bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "None",
			args: args{
				clue:     1,
				variants: [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
				},
				reverseVar: false,
			},
			want: [][]int{},
		},
		{
			name: "Simple one findings",
			args: args{
				clue:     1,
				variants: [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
				},
				reverseVar: false,
			},
			want: [][]int{
				{4, 1, 2, 3},
			},
		},
		{
			name: "Simple several findings",
			args: args{
				clue:     2,
				variants: [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
					{3, 1, 2, 4},
					{4, 3, 2, 1},
					{4, 1, 3, 2},
					{3, 1, 4, 2},
					{1, 4, 3, 2},
					{2, 1, 4, 3},
				},
				reverseVar: false,
			},
			want: [][]int{
				{3, 4, 1, 2},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{1, 4, 3, 2},
				{2, 1, 4, 3},
			},
		},
		{
			name: "Simple several reversed findings",
			args: args{
				clue:     2,
				variants: [][]int{
					{1, 2, 3, 4},
					{2, 3, 4, 1},
					{3, 4, 1, 2},
					{4, 1, 2, 3},
					{3, 1, 2, 4},
					{4, 3, 2, 1},
					{4, 1, 3, 2},
					{3, 1, 4, 2},
					{1, 4, 3, 2},
					{2, 1, 4, 3},
				},
				reverseVar: true,
			},
			want: [][]int{
				{2, 3, 4, 1},
				{3, 4, 1, 2},
				{4, 1, 2, 3},
				{3, 1, 4, 2},
				{2, 1, 4, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterVarsByClue(tt.args.clue, tt.args.variants, tt.args.reverseVar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterVarsByClue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getExistsNums(t *testing.T) {
	type args struct {
		town [][]int
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "None",
			args: args{
				town:    [][]int{
					{0, 0, 2, 3},
					{1, 0, 3, 0},
					{0, 0, 0, 0},
					{0, 0, 2, 4},
				},
				x: 2,
				y: 1,
			},
			want: []int{},
		},
		{
			name: "Simple test",
			args: args{
				town:    [][]int{
					{0, 0, 2, 3},
					{1, 0, 3, 0},
					{0, 0, 0, 0},
					{0, 0, 2, 4},
				},
				x: 3,
				y: 2,
			},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExistsNums(tt.args.town, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExistsNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLineClues(t *testing.T) {
	type args struct {
		clues []int
		n     int
	}
	tests := []struct {
		name  string
		args  args
		wantS int
		wantE int
	}{
		{
			name: "Simple 1",
			args: args{
				clues: []int{0,0,1,2,0,2,0,0,0,3,0,0,0,1,0,0},
				n:     2,
			},
			wantS: 1,
			wantE: 3,
		},
		{
			name: "Simple 2",
			args: args{
				clues: []int{0,0,1,2,0,2,0,0,0,3,0,0,0,1,0,0},
				n:     5,
			},
			wantS: 0,
			wantE: 2,
		},
		{
			name: "Simple 3",
			args: args{
				clues: []int{0,0,1,2,0,2,0,0,0,3,0,0,0,1,0,0},
				n:     0,
			},
			wantS: 0,
			wantE: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotE := getClues(tt.args.clues, tt.args.n)
			if gotS != tt.wantS {
				t.Errorf("getClues() gotS = %v, want %v", gotS, tt.wantS)
			}
			if gotE != tt.wantE {
				t.Errorf("getClues() gotE = %v, want %v", gotE, tt.wantE)
			}
		})
	}
}

func Test_getPosses(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "Simple 1",
			args: args{0},
			want: [][2]int{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
			},
		},
		{
			name: "Simple 2",
			args: args{2},
			want: [][2]int{
				{0, 2},
				{1, 2},
				{2, 2},
				{3, 2},
			},
		},
		{
			name: "Simple 3",
			args: args{4},
			want: [][2]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
		},
		{
			name: "Simple 4",
			args: args{7},
			want: [][2]int{
				{3, 0},
				{3, 1},
				{3, 2},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPosses(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPosses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setOnTown(t *testing.T) {
	type args struct {
		town   [][]int
		posses [][2]int
		sks    []int
	}
	tests := []struct {
		name string
		args args
		wantS [][]int
	}{
		{
			name: "Simple 1",
			args: args{
				town:   [][]int{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posses: [][2]int{
					{0, 0},
					{1, 0},
					{2, 0},
					{3, 0},
				},
				sks:    []int{1, 2, 3, 4},
			},
			wantS: [][]int{
				{1, 0, 0, 0},
				{2, 0, 0, 0},
				{3, 0, 0, 0},
				{4, 0, 0, 0},
			},
		},
		{
			name: "Not enough data 1",
			args: args{
				town:   [][]int{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posses: [][2]int{
					{0, 0},
					{1, 0},
				},
				sks:    []int{1, 2, 3, 4},
			},
			wantS: [][]int{
				{1, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Not enough data 2",
			args: args{
				town:   [][]int{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posses: [][2]int{
					{0, 0},
					{1, 0},
					{2, 0},
					{3, 0},
				},
				sks:    []int{1, 2, 3},
			},
			wantS: [][]int{
				{1, 0, 0, 0},
				{2, 0, 0, 0},
				{3, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Skip set via zero in sks",
			args: args{
				town:   [][]int{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posses: [][2]int{
					{0, 0},
					{1, 0},
					{2, 0},
					{3, 0},
				},
				sks:    []int{1, 0, 2, 3},
			},
			wantS: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{3, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setOnTown(tt.args.town, tt.args.posses, tt.args.sks)
			if !reflect.DeepEqual(tt.args.town, tt.wantS) {
				t.Errorf("setOnTown() = %v, want %v", tt.args.town, tt.wantS)
			}
		})
	}
}

func Test_isFilled(t *testing.T) {
	type args struct {
		town [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple False",
			args: args{[][]int{
				{1, 2, 3, 4},
				{1, 2, 0, 4},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			}},
			want: false,
		},
		{
			name: "Simple True",
			args: args{[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFilled(tt.args.town); got != tt.want {
				t.Errorf("isFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}