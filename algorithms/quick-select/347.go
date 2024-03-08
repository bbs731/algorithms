package quick_select


/****
两种写法，
1.  cnts 用 -cnts 改变大小关系
2.  求 n-k instead of k

3. 注意 j == k 的时候 break,  在loop 外面， 处理 a[k]
 */

func topKFrequent(b []int, k int) []int {
	m := make(map[int]int)
	for _, num := range b {
		m[num]++
	}
	type pair struct {
		cnts int
		val  int
	}

	ll := make([]pair, 0, len(m))

	for num, v := range m {
		ll = append(ll, pair{-v, num})
	}

	l, r := 0, len(ll)-1

	for l < r {
		i, j := l, r+1
		v := ll[l]
		for {
			for i++; i < r && ll[i].cnts < v.cnts; i++ {
			}
			for j--; j > l && ll[j].cnts > v.cnts; j-- {
			}
			if i >= j {
				break
			}
			ll[i], ll[j] = ll[j], ll[i]
		}

		ll[l], ll[j] = ll[j], ll[l]
		if j == k {
			// 这里， 不是唯一的， loop break point, l >= r 的条件也是可以的。 所以这里 break 在外面处理 a[k[
			break
		} else if j > k {
			r = j - 1
		} else {
			l = j + 1
		}
	}

	ans := []int{}
	for p := 0; p < k; p++ {
		ans = append(ans, ll[p].val)
	}
	return ans
}




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

