package weekly


/*
前后缀分解的模版题目

*/


/*
空间优化：改进版本, 不需要创建额外的 prefix, reverse 数组。 所有操作都作用在 ans 数组上。
 */
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	//for i:=0; i <n; i++ {
	//	ans[i]= 1
	//}

	suffix :=1
	for i := n-1; i>=0; i-- {
		ans[i] = suffix
		suffix *= nums[i]
	}

	prefix := 1
	for i:=0; i<n; i++ {
		ans[i] *=prefix
		prefix *= nums[i]
	}
	return ans
}

func productExceptSelf(nums []int) []int {
	n := len(nums)

	ans := make([]int, n)

	//prepare two array  前缀积， 和后缀积
	prefix := make([]int, n+1)
	reverse := make([]int, n+1)

	prefix[0] = 1
	reverse[n] = 1
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		reverse[i] = reverse[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		ans[i] = prefix[i] * reverse[i+1]
	}
	return ans
}
