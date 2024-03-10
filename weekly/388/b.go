package  weekly

import "sort"

/****

[2,3,4,5]
2

 */


func maximumHappinessSum(happiness []int, k int) int64 {
	//for i:=0; i<len(happiness); i++ {
	//	happiness[i] = max(0, happiness[i]-k+1)
	//}
	h := happiness
	sort.Ints(h)
	ans := 0
	cnt := 0
	for i:= len(h)-1; i>= len(h)-k && i>=0; i-- {
		if h[i] >= cnt {
			cnt++
			ans += happiness[i]
		} else {
			break
		}
	}
	ans = ans - cnt*(cnt-1)/2

	return int64(ans)
}