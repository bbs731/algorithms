package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_cf1738B(t *testing.T) {
	testCases := [][2]string{
		{
			`1
4 2
1 3
`,
		`Yes	

`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF1738B)
}
