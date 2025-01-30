// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc163/tasks/abc163_f
// 提交：https://atcoder.jp/contests/abc163/submit?taskScreenName=abc163_f
// 对拍：https://atcoder.jp/contests/abc163/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc163_f&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc163/submissions?f.Status=AC&f.Task=abc163_f&orderBy=source_length
func Test_f(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 2 1
1 2
2 3`,
			`5
4
0`,
		},
		{
			`1
1`,
			`1`,
		},
		{
			`2
1 2
1 2`,
			`2
2`,
		},
		{
			`5
1 2 3 4 5
1 2
2 3
3 4
3 5`,
			`5
8
10
5
5`,
		},
		{
			`8
2 7 2 5 4 1 7 5
3 1
1 2
2 7
4 5
5 6
6 8
7 8`,
			`18
15
0
14
23
0
23
0`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
