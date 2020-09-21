package main

import "fmt"

func main() {
	res := MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Printf("%v\n", res)
}

func MaximumSubarraySum(numbers []int) int {
	max := 0
	for i := 0; i < len(numbers); i++ {
		for _, inc := range []int{-1, 1} {
			if sum := maxByDirection(i, inc, numbers); sum > max {
				max = sum
			}
		}
	}
	return max
}

func maxByDirection(i int, inc int, numbers []int) int {
	max, sum := 0, numbers[i]
	for j := i + inc; j >= 0 && j < len(numbers); j += inc {
		sum += numbers[j]
		if max < sum {
			max = sum
		}
	}
	return max
}
