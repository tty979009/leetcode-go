package test

import (
	"fmt"
	"leetcode-go/leetcode"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "aabaab!bb"
	fmt.Println(leetcode.LengthOfLongestSubstring(s))
}
