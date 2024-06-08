package gapinprimes

import (
	"math"
)

/*
	Challenge link: https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4/train/go

	The prime numbers are not regularly spaced. For example from 2 to 3 the gap is 1. From 3 to 5 the gap is 2.
	From 7 to 11 it is 4. Between 2 and 50 we have the following pairs of 2-gaps primes: 3-5, 5-7, 11-13, 17-19, 29-31, 41-43
	A prime gap of length n is a run of n-1 consecutive composite numbers between two successive prime

	g (integer >= 2) which indicates the gap we are looking for
	m (integer > 2) which gives the start of the search (m inclusive)
	n (integer >= m) which gives the end of the search (n inclusive)

	In the example above gap(2, 3, 50) will return [3, 5]
	which is the first pair between 3 and 50 with a 2-gap.

	So this function should return the first pair of two prime numbers spaced with a gap of g between the limits m, n
	if these numbers exist otherwise nil
*/

func isPrimeNumber(n int) bool {
	nSqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= nSqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func getPrimeNumbersInGap(m, n int) []int {
	var primeNumbers []int

	for i := m; i <= n; i++ {
		if isPrimeNumber(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}

func getGapFromSlice(g int, numbers []int) []int {
	if len(numbers) < 2 {
		return nil
	}

	x, y := numbers[0], numbers[1]
	if y-x == g {
		return []int{x, y}
	}

	for i := 2; i < len(numbers); i++ {
		x = numbers[i]
		if x-y == g {
			return []int{y, x}
		}
		y = x
	}

	return nil
}

func Gap(g, m, n int) []int {
	primesInGap := getPrimeNumbersInGap(m, n)
	return getGapFromSlice(g, primesInGap)
}
