package math

import "strings"

var mod = int(1e9) + 7

func countAnagrams(s string) int {
	sl := strings.Split(s, " ")

	a, b := 1, 1
	for _, word := range sl {
		cnts := make(map[int32]int)
		for i, c := range word{
			a = a * (i+1) % mod
			cnts[c]++
			b = b *cnts[c]%mod
		}
	}

	//https://oi-wiki.org//math/number-theory/inverse/
	// 需要求  (a/b)%mod 这个除法是有难度的， 需要用到费马小定理的逆元的概念
	// 计算 a/b 时候需要找到,  b 在 mod 下的逆元. 因为 mod 是一个质数， 所以 b^(mod-1) = 1 (mod)
	// 因此  1/b  对应  b ^(mod-2) = 1/b (mod)  得到  a * b^(mod-2) = a/b (mod)
	return a * power(b, mod-2) % mod
}



//快速幂，感觉会写了。
func power(num, k int) int {
	ans := 1
	for k > 0 {  //这里注意了， 是 > 0
		if k & 1 !=0 {
			ans = ans * num % mod
		}
		k = k >> 1
		num = num*num % mod
	}
	return ans % mod
}
