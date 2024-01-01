package weekly

import (
	"fmt"
	"sort"
)

func maximumLength(s string) int {
	intervals := [26][]int{}
	track :=[26]int{}

	track[s[0]-'a'] = 1
	for i:=1; i<len(s);i++{
		if s[i] == s[i-1] {
			track[s[i]-'a']++
		} else {
			if track[s[i]-'a'] != 0 {
				intervals[s[i]-'a'] = append(intervals[s[i]-'a'], track[s[i]-'a'])
			}
			track[s[i]-'a']=1
		}
	}
	ans := -1
	for i:=0; i<26; i++ {
		if track[i] != 0 {
			intervals[i] = append(intervals[i], track[i])
		}
		l := intervals[i]
		sort.Sort(sort.Reverse(sort.IntSlice(l)))
		sum := 0
		for j :=0; j<len(l); j++ {
			sum += l[j]
		}
		if len(l) > 0 && sum >=3 {
			ans = max(ans, l[0]-2)
		}
		if len(l) > 1 && sum >=3 {
			if l[0] > l[1] {
				ans = max(ans, l[1])
			} else {
				// l[0] == l[1]
				ans = max(ans, l[0]-1)
				if len(l) > 2 {
					ans = max(ans, l[2])
				}
			}
		}
	}
	return ans
}