package mono_queue

/*
灵神的， 用前缀和， 然后把前缀和，用 单调队列来优化。 解法看起来更加的清晰。 
 */
func shortestSubarray(nums []int, k int) int {
	q := []int{}  //keep the positive nums's index
	sum := 0
	inf := int(1e9)
	ans := inf

	for i, x := range nums {
		// 入队。
		q = append(q, i)
		sum += x

		//出队
		for len(q) >0 && nums[q[len(q)-1]] < 0 {  // remove 没有用的 neg num element, 把它的值加在前面的那个数上。
		// 我们让 q 里面只保留正数。
			neg := nums[q[len(q)-1]]
			q = q[:len(q)-1] // remove last element

			if len(q) > 0 {
				// accumulate the neg to the last element
				nums[q[len(q)-1]] += neg
			}
		}

		// answer
		if len(q) == 0 {
			sum = 0
		}

		for sum >=k {
			//找到解了。
			ans = min(ans, q[len(q)-1] - q[0] + 1)

			// 但是我们可能有更优的解， 所以可以移动队列左端。
			// pop left
			sum -= nums[q[0]]
			q = q[1:]
		}
	}

	if ans == inf {
		return -1
	}
	return ans
}

/*
[84,-37,32,40,95]
 */

func shortestSubarray_wrong(nums []int, k int) int {
	n := len(nums)
	inf := int(1e9)
	ans := inf
	left := 0
	sum := 0
	for i, x := range nums {
		// 入队
		sum += x
		if sum <= 0 {
			left = i+1
			sum = 0
			continue
		}

		//if sum >=k {
		//	ans = min(ans, i-left + 1)
		//}

		//出队
		for sum >=k && left<n {
			ans = min(ans, i-left+1)
			tmp :=0
			// 从 i 倒着数直到遇到 <0 的数， 或者 sum >=k时， 停止
			for j := i; j > left; j-- {
				if nums[j] < 0 {
					break
				}
				tmp += nums[j]
				if tmp >=k {
					ans = min(ans, i-j+1)
					break
				}
			}
			//drop left

			sum -= nums[left]
			left++
			for left <= i && nums[left] <=0 {
				sum -= nums[left]
				left++
			}
		}
	}
	if ans == inf {
		return -1
	}
	return ans
}
