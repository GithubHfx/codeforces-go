// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`"iloveleetcode"`, `["i","love","leetcode","apples"]`, 
			`true`,
		},
		{
			`"iloveleetcode"`, `["apples","i","love","leetcode"]`, 
			`false`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, isPrefixString, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-253/problems/check-if-string-is-a-prefix-of-array/
