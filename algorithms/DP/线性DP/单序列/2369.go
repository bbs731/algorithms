package dp

/***
https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/
灵神的解法

灵神提示，看到 “划分” 猜到是子问题，可以用DP来解。

 */
func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	// 初始化   // 看着这个初始化，就不是很顺眼啊！我又没有办法把它干掉！
	f[0] = true
	if nums[0] == nums[1] {
		f[2] = true
	}
	for i:=2; i<n; i++ {
		if nums[i]== nums[i-1] && f[i-1] == true{
			f[i+1] = true
		}
		if nums[i] == nums[i-1] && nums[i-1] == nums[i-2] && f[i-2]== true{
			f[i+1] = true
		}
		if nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1 && f[i-2] == true {
			f[i+1] = true
		}
	}
	return f[n]
}
