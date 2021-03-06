package main

import "fmt"

func main() {
	r1 := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(r1)
}

/**
* 问题
* 在一个字符串重寻找没有重复字母的最长子串。
* 2020-09-14
* Example:
* Input: "abcabcbb"
* Output: 3
* Explanation: The answer is "abc", with the length of 3.
* Input: "bbbbb"
* Output: 1
* Explanation: The answer is "b", with the length of 1.
* Input: "pwwkew"
* Output: 3
* Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*
*
*/
// 解法一 位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
func lengthOfLongestSubstring_(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
