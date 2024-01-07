package weekly

import "math/bits"



func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)

	type args struct {
		pos, mask int
		changed bool
	}
	memo := make(map[args]int)

	var dfs func(int, int , bool) int
	dfs = func(i int, mask int, changed bool) (res int) {
		if i == n {
			return 0
		}

		a := args{i, mask, changed}
		if v, ok := memo[a]; ok {
			return v
		}

		bit := 1 << (s[i]- 'a')
		newMask := mask | bit
		// no change bit
		if bits.OnesCount(uint(newMask)) > k {
			res =  dfs(i+1, bit, changed) + 1
		}else {
			res =  dfs(i+1, newMask, changed)
		}

		//change bit
		if !changed {
			// loop the char to be changed
			for j:=0; j<26; j++{
				newMask := mask | 1<< j
				if bits.OnesCount(uint(newMask)) > k {
					res = max(res, dfs(i+1, 1 <<j, true)+1)
				}else {
					res = max(res, dfs(i+1, newMask, true))
				}
			}
		}

		memo[a]= res
		return res
	}

	return dfs(0,0, false)
}
