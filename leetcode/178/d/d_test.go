// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`[[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]`}, {`[[1,1,3],[3,2,2],[1,1,4]]`}, {`[[1,2],[4,3]]`}, {`[[2,2,2],[2,2,2]]`}, {`[[4]]`}}
	exampleOuts := [][]string{{`3`}, {`0`}, {`1`}, {`3`}, {`0`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	exampleIns = append(exampleIns, []string{`[[1,1,3],[2,2,2],[4,4,1]]`})
	exampleOuts = append(exampleOuts, []string{`1`})
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, minCost, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-178/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/
//[[1,1,3],[2,2,2],[4,4,1]]