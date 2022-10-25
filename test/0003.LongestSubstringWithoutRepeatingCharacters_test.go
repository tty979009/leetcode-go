package test

import (
	"fmt"
	"leetcode-go/leetcode"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "aabaab!bb"
	fmt.Println(leetcode.LengthOfLongestSubstringM(s))
	fmt.Println(leetcode.LengthOfLongestSubstring(s))
	fmt.Println(leetcode.LengthOfLongestSubstring1(s))
	fmt.Println(leetcode.LengthOfLongestSubstring2(s))
}
