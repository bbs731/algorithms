package quick_select

import (
	"bytes"
	"container/heap"
)

/***

给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。

返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。



示例 1:

输入: s = "aab"
输出: "aba"
示例 2:

输入: s = "aaab"
输出: ""


提示:

1 <= s.length <= 500
s 只包含小写字母

 */

type pair struct {
	c    byte
	cnts int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

// 最大堆
func (h hp) Less(i, j int) bool {
	return h[i].cnts > h[j].cnts
}

func reorganizeString(s string) string {
	tot := len(s)
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	h := make(hp, 0)
	heap.Init(&h)

	for k, v := range m {
		heap.Push(&h, pair{k, v})
	}

	ans := bytes.Buffer{}

	/***
	好题啊， 还是没想明白，最后如何做拼接。
	 */
	for h.Len() > 1 {

	}

}
