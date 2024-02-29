package quick_select

import "math/rand"

/***
给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。



示例 1：

输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
    注意，按字母顺序 "i" 在 "love" 之前。
示例 2：

输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]
解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
    出现次数依次为 4, 3, 2 和 1 次。
 */

/***
经验教训， 这是一个假的， top K 问题， 最会的答案，遵循严格的顺序是唯一的。（不像 top k 问题)
所以， 需要的是 sort, 要不用 heap sort,  要不像问题， 用 quick_selection partition 做 quicksort


ToDo: 可以用 heap 来做一下练习啊！
*/

type pair struct {
	word string
	freq int
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}

	a := make([]pair, 0, len(m))
	for k, v := range m {
		a = append(a, pair{k, v})
	}

	sortArray(a)

	ans := make([]string, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, a[i].word)
	}

	return ans

}

func partition(a []pair, k int) {
	n := len(a)
	rand.Shuffle(n, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	// 写 partition 的方式
	l, r := 0, len(a)-1
	for l < r {
		i, j := l, r+1
		v := a[l]
		for {
			for i++; i < r && (a[i].freq > v.freq || (a[i].freq == v.freq && a[i].word < v.word)); i++ {
			}
			for j--; j > l && (a[j].freq < v.freq || (a[j].freq == v.freq && a[j].word > v.word)); j-- {
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
	return
}

func sortArray(a []pair) {
	n := len(a)
	if n <= 1 {
		return
	}

	mid := n >> 1
	partition(a, mid)
	sortArray(a[:mid])
	sortArray(a[mid+1:])
}
