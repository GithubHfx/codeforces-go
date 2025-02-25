// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2009/G2
// https://codeforces.com/problemset/status/2009/problem/G2?friends=on
func Test_cf2009G2(t *testing.T) {
	testCases := [][2]string{
		{
			`3
7 5 3
1 2 3 2 1 2 3
1 7
2 7
3 7
8 4 2
4 3 1 1 2 4 3 2
3 6
1 5
5 4 2
4 5 1 2 3
1 4
1 5`,
			`6
5
2
2
5
2
3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2009G2)
}
