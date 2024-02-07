package bits_operation

/***
根据 2411 的模版写的
 */
func subarrayBitwiseORs(nums []int) int {
	n := len(nums)
	ans := make(map[int]struct{})
	type pair struct{ or, i int }
	ors := []pair{} // 按位或的值 + 对应子数组的右端点的最小值
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		ors = append(ors, pair{0, i})
		ors[0].or |= num
		ans[ors[0].or]= struct{}{}
		k := 0
		for _, p := range ors[1:] {
			p.or |= num
			if ors[k].or == p.or {
				ors[k].i = p.i // 合并相同值，下标取最小的
			} else {
				k++
				ors[k] = p
				ans[p.or] = struct{}{}
			}
		}
		ors = ors[:k+1]
	}
	return len(ans)
}