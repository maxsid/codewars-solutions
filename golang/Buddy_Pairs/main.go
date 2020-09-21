package main

import "math"

func main() {

}

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {
	for n := start; n < limit; n++ {
		m := getSumOfDivisorsSqrt(n) - 1
		if n2 := getSumOfDivisorsSqrt(m) - 1; n2 == n {
			return []int{n, m}
		}
	}
	return nil
}

func getSumOfDivisors(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func getSumOfDivisorsSqrt(n int) int {
	sum := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i < sqrtN; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
