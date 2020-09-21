package main

import (
	"reflect"
	"testing"
)

func TestBuddy(t *testing.T) {
	type args struct {
		start int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Easy 1",
			args: args{10, 50},
			want: []int{48, 75},
		},
		{
			name: "Easy 2",
			args: args{48, 50},
			want: []int{48, 75},
		},
		// {
		// 	name: "should handle basic cases buddy 1",
		// 	args: args{1071625, 1103735},
		// 	want: []int{1081184, 1331967},
		// },
		// {
		// 	name: "should handle basic cases buddy 2",
		// 	args: args{57345, 90061},
		// 	want: []int{62744, 75495},
		// },
		// {
		// 	name: "should handle basic cases buddy 3",
		// 	args: args{2693, 7098},
		// 	want: []int{5775, 6128},
		// },
		// {
		// 	name: "should handle basic cases buddy 1",
		// 	args: args{6379, 8275},
		// 	want: nil,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Buddy(tt.args.start, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buddy() = %v, want %v", got, tt.want)
			}
		})
	}
}
