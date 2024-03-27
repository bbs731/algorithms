package one_day_exercise

import (
	"sort"
)

func countWays(ranges [][]int) int {

	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	if len(ranges) == 0 {
		return 0
	}

	cnts := 1
	end := ranges[0][1]

	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= end {
			end = max(end, ranges[i][1])
			continue
		}
		cnts++
		end = ranges[i][1]
	}
	//fmt.Println(cnts)
	return power(2, cnts)
}

func power(num, p int) int {
	res := 1
	MOD := int(1e9 + 7)

	for i := 1; i <= p; i++ {
		res = res * 2 % MOD
	}
	return res

	// 这个是教训啊， 快速幂，不适合取模啊！
	// for p > 0 {
	// 	if p&1 > 0 {
	// 		res = res * num
	// 		res %= MOD
	// 	} else {
	// 		num = num * num
	// 		num %= MOD
	// 	}
	// 	p = p >> 1
	// }
	// return res % MOD
}
