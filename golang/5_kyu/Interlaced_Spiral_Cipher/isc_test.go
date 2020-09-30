package isc

import (
	"reflect"
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encode 1",
			args: args{"Romani ite domum"},
			want: "Rntodomiimuea  m",
		},
		{
			name: "Encode 2",
			args: args{"Sic transit gloria mundi"},
			want: "Stsgiriuar i ninmd l otac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.s); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encode 1",
			args: args{"Rntodomiimuea  m"},
			want: "Romani ite domum",
		},
		{
			name: "Encode 2",
			args: args{"Stsgiriuar i ninmd l otac"},
			want: "Sic transit gloria mundi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.s); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNextPosition(t *testing.T) {
	type args struct {
		pos        [2]int
		dirN       int
		squareSize int
		step       int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 [2]int
	}{
		{
			name:  "From 0;0 to 0;4",
			args:  args{[2]int{0, 0}, 0, 5, 4},
			want:  0,
			want1: [2]int{0, 4},
		},
		{
			name:  "From 0;1 to 1;4",
			args:  args{[2]int{0, 1}, 0, 5, 4},
			want:  1,
			want1: [2]int{1, 4},
		},
		{
			name:  "From 0;2 to 2;4",
			args:  args{[2]int{0, 2}, 0, 5, 4},
			want:  1,
			want1: [2]int{2, 4},
		},
		{
			name:  "From 3;4 to 4;1",
			args:  args{[2]int{3, 4}, 1, 5, 4},
			want:  2,
			want1: [2]int{4, 1},
		},
		{
			name:  "From 4;1 to 1;0",
			args:  args{[2]int{4, 1}, 2, 5, 4},
			want:  3,
			want1: [2]int{1, 0},
		},
		{
			name:  "From 1;0 to 0;3",
			args:  args{[2]int{1, 0}, 3, 5, 4},
			want:  0,
			want1: [2]int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getNextPosition(tt.args.pos, tt.args.dirN, tt.args.squareSize, tt.args.step)
			if got != tt.want {
				t.Errorf("move() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("move() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
