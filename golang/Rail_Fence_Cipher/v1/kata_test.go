package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Kata 1",
			args: args{"WEAREDISCOVEREDFLEEATONCE",3},
			want: "WECRLTEERDSOEEFEAOCAIVDEN",
		},
		{
			name: "Kata 2",
			args: args{"Hello, World!",3},
			want: "Hoo!el,Wrdl l",
		},
		{
			name: "Empty",
			args: args{"",3},
			want: "",
		},
		{
			name: "The same result",
			args: args{"ABCD",4},
			want: "ABCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Kata 1",
			args: args{"WECRLTEERDSOEEFEAOCAIVDEN",3},
			want: "WEAREDISCOVEREDFLEEATONCE",
		},
		{
			name: "Kata 2",
			args: args{"Hoo!el,Wrdl l",3},
			want: "Hello, World!",
		},
		{
			name: "Empty",
			args: args{"",3},
			want: "",
		},
		{
			name: "The same result",
			args: args{"ABCD",4},
			want: "ABCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRailsEquality(t *testing.T) {
	resEncode := getRailsFromDecodedStr("WEAREDISCOVEREDFLEEATONCE", 3)
	resDecode := getRailsFromEncodedStr("WECRLTEERDSOEEFEAOCAIVDEN", 3)
	if !reflect.DeepEqual(resDecode, resEncode) {
		t.Errorf("getRailsFromDecodedStr() != getRailsFromEncodedStr(). First: %v, Second: want %v", resEncode, resDecode)
	}
}

func Test_getRails(t *testing.T) {
	originals := []string{"", ""}
	for len(originals[0]) < 25 {
		originals[0] += "A"
		originals[1] += "A"
	}
	originals[1] += "A"
	for _, original := range originals {
		for n := 2; n <= len(original); n++ {
			rails := getRailsFromDecodedStr(original, n)
			fmt.Printf("rails = %d -> ", n)
			for _, r := range rails {
				fmt.Printf("\t%d;", len(r))
			}
			fmt.Println()
		}
	}
}