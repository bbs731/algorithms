package sieveofetratosthenes

/*
非常清楚的讲解，和数学证明
https://www.cnblogs.com/pxsong/p/SieveofEratosthenes.html
 */

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CountPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ans++
		}
	}
	return ans
}

func CounterPrimesSieve(n int) int {
	primes := make([]bool, n)
	for i := 0; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i < n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	ans := 0
	for i := 2; i < n; i++ {
		if primes[i] {
			ans++
		}
	}
	return ans
}
