package backtracking

import (
	"fmt"
	"testing"
)

func TestGenerateParathesis(t *testing.T) {
	for _, s := range generateParenthesis(8) {
		fmt.Println(s)
	}
}
