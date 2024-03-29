package primewithevendigits

import (
	"strconv"
)

// Sieve of Eratosthenes algorithm
func getAllPrimeNumbersFromN(n int) (primeNumbers []int) {
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if prime[p] {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if prime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func hasEvenDigit(n int) (has bool, count int) {
	for _, v := range strconv.Itoa(n) {
		if int(v)%2 == 0 {
			has = true
			count++
		}
	}
	return
}

func F(n int) int {
	var max int
	var maxCount int

	for _, primeNumber := range getAllPrimeNumbersFromN(n) {
		if has, count := hasEvenDigit(primeNumber); has && primeNumber > max && count >= maxCount && primeNumber != n {
			max = primeNumber
			maxCount = count
		}
	}

	return max
}
