package c300

import (
	"testing"
)

type question struct {
	nums []int
	ans  int
}

func TestOk(t *testing.T) {

	inputs := []question{
		{

			nums: []int{7, 8, 9, 1, 2, 3, 4, 5, 6},
			ans:  6,
		},
		{
			nums: []int{},
			ans:  0,
		},
		{
			[]int{0},
			1,
		},
	}

	for _, input := range inputs {
		if lengthOfLIS(input.nums) != input.ans {
			t.Errorf("Test failed for:%v and expected ans is: %d", input.nums, input.ans)
		}
	}
}
