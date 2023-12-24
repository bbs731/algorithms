package weekly

import "fmt"

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	inf := int(1e9)
	dp := make([]int, n)
	for i:=0; i<n; i++ {
		dp[i]= inf
	}

	if source[0] == target[0] {
		dp[0]= 0
	}

	for i:=1; i <n; i++ {
		for j :=0; j<len(original); j++ {
			p :=original[j]
			if source[i] == target[i]{
				dp[i]= dp[i-1]
			}
			// match ?
			if  i-len(p) +1  >=0 &&  source[i-len(p)+1:i+1] == p && target[i-len(p)+1:i+1] == changed[j] {
				if i - len(p) + 1 ==  0 {
					dp[i] = min(dp[i], cost[j])
				} else {
					dp[i] = min(dp[i], dp[i-len(p)]+cost[j])
				}
			}

		}
	}
	fmt.Println(dp)
	if dp[n-1] == inf {
		return -1
	}
	return int64(dp[n-1])
}