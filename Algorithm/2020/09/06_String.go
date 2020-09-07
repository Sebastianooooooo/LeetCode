package main

import (
	"fmt"
	"strconv"
)

func main() {
	r1 := numDecodings("226")
	fmt.Println(r1)
}

/**
* 问题
* 一条包含字母 A-Z 的消息通过以下方式进行了编码：
	'A' -> 1
	'B' -> 2
	...
	'Z' -> 26
* 给定一个只包含数字的非空字符串，请计算解码方法的总数。
* 2020-09-06
* Example:
* Input: "226"
* Output: 3
* Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*
* Input: "12"
* Output: 2
* Explanation: It could be decoded as "AB" (1 2) or "L" (12).
*
*/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[:1] == "0" {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
