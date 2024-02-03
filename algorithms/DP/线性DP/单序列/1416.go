package dp

/***
有调试了半个小时，还是太慢了，蛋蛋！


https://leetcode.cn/problems/filling-bookcase-shelves/description/

感谢灵神，这种类型的 DP 会了。
2369. 检查数组是否存在有效划分
1416. 恢复数组
1043. 分隔数组以得到最大和
1105. 填充书架

 */
func numberOfArrays(s string, k int) int {
	n := len(s)
	kl := 0
	for tmp:=k; tmp !=0; tmp = tmp/10 {
		kl++
	}

	power10 := make([]int, kl+1)
	power10[0] = 1
	for i:=1; i<=kl; i++ {
		power10[i] = 10* power10[i-1]
	}

	f := make([]int, n+1)
	f[0] = 1

	for i:=0; i<n; i++ {
		candidate := 0
		for j:=i; j>=0 && i-j+1 <=kl; j-- {
			//candidate = candidate*10 + int(s[j]-'0')  // 这里太容易出错了， 分清，那个是 MSB 那个是 LSB
			candidate = power10[i-j]*int(s[j]-'0') + candidate
			if s[j] != '0' && candidate <=k && candidate >0  {
				f[i+1] = (f[i+1] + f[j])% int(1e9 + 7)
			}
		}
	}
	return f[n]
}
