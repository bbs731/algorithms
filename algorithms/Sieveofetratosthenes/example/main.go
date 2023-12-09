package main

import (
	"Sieveofetratosthenes"
	"fmt"
	"time"
)

func main() {
	var start time.Time
	var end time.Duration

	start = time.Now() // 开始时间
	fmt.Println("CountPrimes")
	fmt.Printf("answer is : %d\n", Sieveofetratosthenes.CountPrimes(10000000))
	end = time.Since(start)
	fmt.Printf("total time: %s\n", end)

	start = time.Now() // 开始时间
	fmt.Println("CountPrimesSieve")
	fmt.Printf("answer is : %d\n", Sieveofetratosthenes.CounterPrimesSieve(10000000))
	end = time.Since(start)
	fmt.Printf("total time: %s\n", end)

}
