package Sieveofetratosthenes

/*
非常清楚的讲解，和数学证明
https://www.cnblogs.com/pxsong/p/SieveofEratosthenes.html

1 + 1/2 + 1/3 + 1/4 + ..... 1/n  = O(ln(n))

1/2 + 1/3 + 1/5 + 1/7 + 1/11 + ..... + 1/p  其中 p 为质数。 质数的分布函数是欧拉乘积公式给的。
可以这样理解吗？ 质数在自然数中是近似按照 log 函数分布的

Sum(1/p) = ln( sum(1/n)) = ln(ln(n))

所以 SieveOfEratosthenes 的时间复杂度是：
n*(Sum(1/p)) = n *ln(ln(n))

 */

func CounterPrimesSieve() []bool {

	n := int(1e5) + 1
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	return primes
}
