package interview_questions

/****
面试，你还是差了很多啊， 咋办啊？ 这种智商的题目都做不过， 哎！
 */
func singleNumber(nums []int) int {
	ans := int32(0) // 就是为了，能处理 ans 可能是负数的情况
	for i:=0; i< 32; i++ {
		cnts := 0
		for _, x := range nums {
			if x & (1 << i) !=0 {  // 这个 （） 非常的重要， &  和  << 的优先级问题， 会超出你的预想。
				cnts++
			}
		}
		if cnts%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
