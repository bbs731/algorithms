package dp

/***
太简单了， 但是还是想了半天， 能又半个小时？ 嗯，距离上一次 commit 28 分钟， 太慢了!
 */
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	f := make([]int,n+1)
	for i, num := range arr {
		maxnum := num
		for j:=i; j>=0 && i-j+1<=k; j-- {
			maxnum = max(maxnum, arr[j])
			f[i+1] = max(f[i+1], maxnum *(i-j+1) + f[j])
		}
	}
	return f[n]
}
