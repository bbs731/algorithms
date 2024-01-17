package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1909/C
// https://codeforces.com/problemset/status/1909/problem/C
func Test_cf1909C(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2
8 3
12 23
100 100
4
20 1 2 5
30 4 3 10
2 3 2 3
10
76042 155685 62534 162770 779 3495 97453 122787 86743 142857
81292 172088 146526 199306 36432 165338 168285 127772 119677 151891
1 8 2 9 1 1 1 4 7 4
`,
			`2400
42
609373`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1909C)
}

