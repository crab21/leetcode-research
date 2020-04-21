package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalind(t *testing.T) {
	palind := LongestPalind("abccba")
	fmt.Println(palind)
}

/**
input:abccba
output:
[# a # b # c # c # b # a]
6
6

input:abcdcba
output:
[# a # b # c # d # c # b # a]
7
7
 */
func TestPalindrome(t *testing.T) {
	Palindrome("abccba")
	Palindrome("abcdcba")
}