package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_bMode(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 20
5 199
0 1000000000`, `21
233
2553052375
`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cfbmode2)
}
