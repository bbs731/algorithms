package quick_select

/***

这道题，top 最大 k  的问题 等价于  top n-k 小的问题， 从 n-k 到 n-1 这部分就是 top k 最大的了。
具体看 347_1.go 的题解。
 */

func topKFrequent(b []int, k int) []int {
	m := make(map[int]int)
	for _, num := range b {
		m[num]++
	}
	type pair struct{ v, i int }
	a := make([]pair, 0, len(m))
	for num, v := range m {
		a = append(a, pair{v, num})
	}

	n := len(a)
	l, r := 0, n-1

	for l < r {
		i, j := l, r+1
		v := a[l].v
		for {
			// 注意， 这里的大小关系， 换了， 因为求得是从大到小的顺序。
			for i++; i < r && a[i].v > v; i++ {
			}
			for j--; j > l && a[j].v < v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], a[l]
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}

	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, a[i].i)
	}
	return ans
}
