package main

import (
	"fmt"
	"sort"
)

/*
实现方式，参考：  https://oi-wiki.org/string/sa/
时间复杂度: O(n*lgn*lgn)
 */
type SA struct {
	src    string
	sa     []int
	rk     []int
	height []int
}

/*

可以这么理解，
代码的主线，就是 rank 的倍增排序， sa 是为 求 rank 来服务的， sa 作为辅助的数组。

总共需要 log(n) 轮的倍增， 除了最后一轮， 其它轮， 任意两个 rank[i], rank[j] 是有可能相等的， 这个时候， 需要借助 sa 数组， 来给新的 rank 赋值（相对大小），
当然如果用基数排序来做，也是为了同样的目的，就是给rank赋相对大小的值

思路就是：
1. 如果  rank[sa[i]] == rank[sa[i]-1]  && rank[sa[i] + w ] == rank[sa[i-1] + w]   // 包括倍增之后的结果一样的话
那么下一轮了的 rank,  newrank[sa[i]]  = newrank[sa[i-1]]  (代表，我们还需要一轮 binary lifting, 才能区分 rank 的值）
否则的话，  newrank[sa[i]] = newranks[sa[i-1]] + 1

2. 根据 newrank 的值， 还需要更新下 sa 数组， 这时候，根据 newrank 值， sort 一下 sa 就可以

 */
func NewSA(src string) *SA {
	n := len(src)

	sa := make([]int, n)
	rk := make([]int, 2*n) // 开2n 的大小，是为了  sa[i] + w 就不用检查越不越界了
	for i := n; i < 2*n; i++ {
		rk[i] = -1 // 为了，处理，譬如 str = "aaaaaaaa" 的情况
	}
	oldrk := make([]int, 2*n)

	for i := 0; i < n; i++ {
		rk[i] = int(src[i] - 'a')
		sa[i] = i
	}

	// binary lifting
	for w := 1; w < n; w = w * 2 {
		sort.Slice(sa, func(x int, y int) bool {
			if rk[sa[x]] != rk[sa[y]] {
				return rk[sa[x]] < rk[sa[y]]
			}
			return rk[sa[x]+w] < rk[sa[y]+w]
		})

		// 计算新的 rk 的时候，会有覆盖发生，所以先 copy 一份
		copy(oldrk, rk)
		rk[sa[0]] = 0 // rank 的值域应该从 0 开始
		for i := 1; i < n; i++ {
			if oldrk[sa[i]] == oldrk[sa[i-1]] && oldrk[sa[i]+w] == oldrk[sa[i-1]+w] {
				rk[sa[i]] = rk[sa[i-1]]
				//rk[sa[i]] = p
			} else {
				//p++
				//rk[sa[i]] = p
				rk[sa[i]] = rk[sa[i-1]] + 1
			}
		}
	}

	// 根据引理:  height[rk[i]] >= height[rk[i-1]] - 1.   证明在这里： https://oi-wiki.org/string/sa/
	/*
		推倒的过程，
		根据定义：  height[i] = lcp( suffix(sa[i]), suffix(sa[i-1])
		所以：
			height[rk[i]] = lcp (suffix(sa[rk[i]]),  suffix(sa[rk[i]-1))
	                      = lcp (suffix(i), suffix(sa[rk[i]-1]))

		令， height[rk[i-1]] = h,   则  height[rk[i]] >= h -1
	    令， j = sa[rk[i]-1]  则，

	     height[rk[i]] = lcp(suffix(i), suffix(j))
	                   = h-1 + lcp(suffix(i+h-1），suffix(j+h-1))
		为以下代码，计算 height[rk[i]] 的步骤过程
	 */

	rk[sa[0]] = 0 // 这行代码，可以防止在 n= 1 的时候，程序崩溃， 因为 n=1 时候， rk[0] = src[0] - 'a'  可能会引起在计算 height 数组的时候， sa[rk[i]-1] index out of range
	// compute height
	height := make([]int, n)
	h := 0
	for i := 0; i < n; i++ {
		if rk[i] == 0 {
			continue
		}
		if h > 0 {
			h--
		}
		for j := sa[rk[i]-1]; i+h < n && j+h < n && src[i+h] == src[j+h]; h++ {
		}
		height[rk[i]] = h
	}

	return &SA{src, sa, rk, height}
}

func main() {

	str := "p"
	sa := NewSA(str)
	fmt.Println(sa.sa)
	fmt.Println(sa.rk)
	fmt.Println(sa.height)
	for i := 0; i < len(str); i++ {
		fmt.Println(str[sa.sa[i]:])
	}
}
