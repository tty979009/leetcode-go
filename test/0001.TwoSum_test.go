package test

import (
	"fmt"
	"leetcode-go/leetcode"
	"testing"
)

func TestXxx(t *testing.T) {
	// nums := []int{2, 7, 11, 15}
	nums := []int{11, 7, 2, 15}
	target := 9
	result := leetcode.TwoSumM(nums, target)
	fmt.Println(result)
}
