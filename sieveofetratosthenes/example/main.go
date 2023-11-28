package main

import (
	"fmt"
	"sieveofetratosthenes"
	"time"
)

func main() {
	var start time.Time
	var end time.Duration

	start = time.Now() // 开始时间
	fmt.Println("CountPrimes")
	fmt.Printf("answer is : %d\n", sieveofetratosthenes.CountPrimes(10000000))
	end = time.Since(start)
	fmt.Printf("total time: %s\n", end)

	start = time.Now() // 开始时间
	fmt.Println("CountPrimesSieve")
	fmt.Printf("answer is : %d\n", sieveofetratosthenes.CounterPrimesSieve(10000000))
	end = time.Since(start)
	fmt.Printf("total time: %s\n", end)

}
