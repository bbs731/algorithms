package weekly

import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	mod := 1_000_000_007
	hFences  = append(hFences, m)
	hFences = append(hFences,1)
	vFences = append(vFences, n)
	vFences = append(vFences, 1)
	sort.Ints(vFences)
	sort.Ints(hFences)

	hMap := make(map[int]int)
	vMap := make(map[int]int)

	for i:=0; i<len(hFences)-1; i++{
		for j:= i+1; j<len(hFences); j++ {
			hMap[ hFences[j]-hFences[i]]++
		}
	}

	for i:=0; i<len(vFences)-1; i++ {
		for j := i+1; j<len(vFences); j++ {
			vMap[vFences[j]-vFences[i]]++
		}
	}

	hl := make([]int, 0, len(hMap))
	for hk := range hMap {
		hl = append(hl, hk)
	}
	sort.Ints(hl)

	for i:=len(hl)-1; i>=0; i-- {
		if _, ok := vMap[hl[i]]; ok {
			//return hMap[hl[i]] * vMap[hl[i]]
			return hl[i]*hl[i] % mod
		}
	}
	return -1

}