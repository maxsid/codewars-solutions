package main

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "Empty arrays should have a max of 0",
			numbers: []int{},
			want:    0,
		},
		{
			name:    "Example array should have a max of 6",
			numbers: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want:    6,
		},
		{
			name:    "Example array with all negative values should have a max of 0",
			numbers: []int{-2, -1, -3, -4, -1, -2, -1, -5, -4},
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubarraySum(tt.numbers); got != tt.want {
				t.Errorf("MaximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
