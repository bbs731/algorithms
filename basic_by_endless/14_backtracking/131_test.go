package backtracking

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	var ans [][]string

	ans = partition2("aabc")
	fmt.Println(ans)

	ans = partition("aabc")
	fmt.Println(ans)

}
