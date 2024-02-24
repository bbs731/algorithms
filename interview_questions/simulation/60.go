package simulation

import "fmt"


/***
根 440 这套题， 非常的像

都是在找 kth element 的题目。 找到顺序的规律！
 */

func getPermutation(n int, k int) string {
	fac := make([]int, n+1)
	fac[1] = 1
	for i:=2; i<=n; i++ {
		fac[i] = i*fac[i-1]
	}

	ans := []byte{}
	cur := 1
	mask := 0
	l := n-1

	for k >0 && l>0 {
		if k > fac[l] {  // 这个 > 不是 >= 注意了，这里容易bug
			for cur++; mask & (1<<cur) !=0; cur++{}
			k -= fac[l]
		} else {
			// select cur
			ans = append(ans, byte(cur+'0'))
			mask |= 1 << cur

			for i:=1; i<=n; i++ {
				if mask& (1<<i) !=0  {
					continue
				}
				cur = i
				break
			}
			l--
		}
	}
	// 我们还差了最后一位， 真正反正遍历，都没问题，找到哪一位，补上
	//for i:=n; i>=1; i-- {
	for i:=1; i<=n; i++ {
		if mask & (1<<i) == 0{
			ans = append(ans, byte(i+'0'))
		}
	}
	return string(ans)
}