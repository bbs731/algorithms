package Sieveofetratosthenes

import (
	"fmt"
	"testing"
)

const (
	N = 10000000
)

func BenchmarkCountPrimes(b *testing.B) {
	var ans int
	for i := 0; i < b.N; i++ {
		ans = CountPrimes(N)
	}

	fmt.Println(ans)
}

func BenchmarkCountPrimesSieve(b *testing.B) {
	var ans int

	for i := 0; i < b.N; i++ {
		ans = CounterPrimesSieve(N)
	}
	fmt.Println(ans)
}
