package quick_select

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
	// top k 最大  等价于  top n-k 最小。 从 n-k 开始到  n-1 为最k 大。 这两个问题是等价的
	kprime := n - k

	for l < r {
		i, j := l, r+1
		v := a[l].v
		for {
			for i++; i < r && a[i].v < v; i++ {
			}
			for j--; j > l && a[j].v > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], a[l]
		if j == kprime {
			break
		} else if j < kprime {
			l = j + 1
		} else {
			r = j - 1
		}
	}

	ans := make([]int, 0, k)
	for i := kprime; i < n; i++ {
		ans = append(ans, a[i].i)
	}
	return ans
}
