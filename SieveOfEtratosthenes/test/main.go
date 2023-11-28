package main

import (
	"fmt"
	"github.com/bbs731/Algorithms/SieveOfEtratosthenes"
	"time"
)

//func isPrime(n int) bool {
//	for i := 2; i*i <= n; i++ {
//		if n%i == 0 {
//			return false
//		}
//	}
//	return true
//}
//
//func CountPrimes(n int) int {
//	ans := 0
//	for i := 2; i < n; i++ {
//		if isPrime(i) {
//			ans++
//		}
//	}
//	return ans
//}

func main() {
	start := time.Now() // 开始时间
	fmt.Println("CountPrimes")
	fmt.Printf("answer is : %d\n", SieveOfEtratosthenes.CountPrimes(10000000))
	end := time.Since(start)
	fmt.Printf("total time: %s\n", end)

	start = time.Now() // 开始时间
	fmt.Println("CountPrimesSieve")
	fmt.Printf("answer is : %d\n", SieveOfEtratosthenes.CounterPrimesSieve(10000000))
	end := time.Since(start)
	fmt.Printf("total time: %s\n", end)

}
