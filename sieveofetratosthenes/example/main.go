package main

import (
	"fmt"
	"github.com/bbs731/algorithms/sieveofetratosthenes"
	"time"
)


func main() {
	start := time.Now() // 开始时间
	fmt.Println("CountPrimes")
	fmt.Printf("answer is : %d\n", sieveOfEtratosthenes.CountPrimes(10000000))
	end := time.Since(start)
	fmt.Printf("total time: %s\n", end)

	start = time.Now() // 开始时间
	fmt.Println("CountPrimesSieve")
	fmt.Printf("answer is : %d\n", sieveOfEtratosthenes.CounterPrimesSieve(10000000))
	end := time.Since(start)
	fmt.Printf("total time: %s\n", end)

}
