package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/**
最大子序列的和问题
动态规划问题
*/
func MaxLengthLengthOfInner(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for i := 0; i < len(nums); i++ {
		if curSum < 0 {
			curSum = 0
		}
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

/**
有一座高度是10级台阶的楼梯，从下往上走，每跨一步只能向上1级或者2级台阶。要求用程序来求出一共有多少种走法

分类：动态规划问题
*/
//  1-1递归做法
// time: O(2^n)   效率很低
func getClimbingStep(n int) int {
	if n < 1 {
		return 0
	}
	if n < 2 {
		return 1
	}
	if n < 3 {
		return 2
	}
	return getClimbingStep(n-1) + getClimbingStep(n-2)
}

func getClimbingStepByDP(n int) int {
	if n < 1 {
		return 0
	}
	if n < 2 {
		return 1
	}
	if n < 3 {
		return 2
	}
	a, b, tmp := 1, 2, 0
	for i := 3; i <= n; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}

/**
最长回文子串  中心扩散法
*/
func Palindrome(str string) {
	if str == "" || len(str) == 0 {
		return
	}
	var sbStr = []string{}
	for k := range str {
		u := str[k]
		sbStr = append(sbStr, "#")
		sbStr = append(sbStr, string(u))
	}
	sbStr = append(sbStr, "#")
	var max_length = 0
	for k := range sbStr {
		kp := subpalidromelen(sbStr, k)
		if kp > 1 {
			skp := kp / 2
			//防止kp为奇数
			skpRest := kp %2
			fmt.Println(sbStr[k-skp*2-skpRest:k+skp*2+skpRest])
			fmt.Println(kp)
		}
		if kp > max_length {
			max_length = kp
		}
	}
	fmt.Println(max_length)

}

func subpalidromelen(str []string, i int) int {
	var lens = 0
	for k := 0; k <= i && k < len(str)-i; k++ {
		if str[i-k] == str[i+k] {
			lens++
		} else {
			break
		}
	}
	return lens - 1
}

/**
中心展开法-2
时间复杂度：O(n^2) 空间复杂度：O(1)
*/
func LongestPalind(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}

	var sp_result = ""
	for i := 0; i < len(s); i++ {
		sp_result = helper(s, i, i, sp_result)
		sp_result = helper(s, i, i+1, sp_result)
	}
	return sp_result
}

func helper(s string, left int, right int, sp_result string) string {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; {
		left--
		right++
	}
	var cur = s[left+1 : right]
	if len(cur) > len(sp_result) {
		sp_result = cur
	}
	return sp_result
}

/**
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/
func numDecodings(s string) int {
	if s == "" || len(s) == 0 || strings.Index(s, "0") == 0 {
		return 0
	}
	var dp []int
	dp = append(dp, 1, 1)
	for i := 2; i <= len(s); i++ {
		first, eer := strconv.Atoi(s[i-2 : i])
		if eer != nil {
			continue
		}
		firstStep := 0
		if first >= 10 && first <= 26 {
			firstStep = dp[i-2]
		}
		second, eer := strconv.Atoi(s[i-1 : i])
		if eer != nil {
			continue
		}

		if second > 0 && second < 10 {
			firstStep += dp[i-1]
		}
		dp = append(dp, firstStep)
	}
	return dp[len(s)]
}

/**
A message containing letters from A-Z is being encoded to numbers using the following mapping way:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Beyond that, now the encoded string can also contain the character '*', which can be treated as one of the numbers from 1 to 9.

Given the encoded message containing digits and the character '*', return the total number of ways to decode it.

Also, since the answer may be very large, you should return the output mod 109 + 7.

Example 1:
Input: "*"
Output: 9
Explanation: The encoded message can be decoded to the string: "A", "B", "C", "D", "E", "F", "G", "H", "I".
Example 2:
Input: "1*"
Output: 9 + 9 = 18
Note:
The length of the input string will fit in range [1, 105].
The input string will only contain the character '*' and digits '0' - '9'.
*/
var mod int = 1e9+7

func numDecodingsII(s string) int {
	if !strings.Contains(s, "*") {
		return numDecodings(s)
	}
	if s == "" || len(s) == 0 || strings.Index(s, "0") == 0 {
		return 0
	}
	var dp = [2]int{1, 1}
	var ans = 0
	var pre uint8 = 0
	for i := 0; i < len(s); i++ {
		cur := s[i]
		pre=0
		if i > 	0 {
			pre = s[i-1]
		}
		if cur == '0' && pre == '0' {
			return 0
		}
		if cur != '*' && pre != '*' && pre>'2'&&cur=='0'{
			return 0
		}

		ans = 0
		if cur == '0' {
			if pre == '*' {
				ans = (dp[0] * 2) % mod
			} else {
				ans = dp[0]
			}
		} else if cur == '*' {
			if pre == '*' {
				ans = (dp[0] * 15) % mod
			} else if pre == '1' {
				ans = (dp[0] * 9) % mod
			} else if pre == '2' {
				ans = dp[0] * 6 % mod
			}
			ans = (ans + dp[1]*9) % mod
		} else if cur >= '1' && cur <= '6' {
			if pre == '*' {
				ans = dp[0] * 2 & mod
			} else if pre >= '1' && pre <= '2' {
				ans = dp[0] % mod
			}
			ans = (ans + dp[1]) % mod
		} else {
			if pre == '*' || pre == '1' {
				ans = dp[0] % mod
			}
			ans = (ans + dp[1]) % mod
		}
		dp[0] = dp[1]
		dp[1] = ans
	}
	return ans
}
func DPStart() {
	nano := time.Now().UnixNano()
	decodings := numDecodingsII("**1**")
	//decodings := numDecodingsII("*******")
	fmt.Println(time.Now().UnixNano()- nano)
	fmt.Println(decodings)
}
