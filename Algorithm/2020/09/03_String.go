package main

import (
	"fmt"
	"strings"
)

func main() {
	r1 := strStr("aaaaa","ba")
	fmt.Println(r1)
	r2 := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(r2)
}


/**
* 问题1
* 实现一个查找 substring 的函数。如果在母串中找到了子串，返回子串在母串中出现的下标，如果没有找到，返回 -1，如果子串是空串，则返回 0 。
* 2020-09-03
* Example:
* Input: haystack = "aaaaa", needle = "bba"
* Output: -1
*
*
*/
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// 解法二
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}


/**
* 问题2
* 判断所给的字符串是否是有效的回文串。<“回文串”是一个正读和反读都一样的字符串>
* 2020-09-03
* Example:
* "A man, a plan, a canal: Panama" is a palindrome.
* "race a car" is not a palindrome.
*
*
 */
func isPalindrome(s string) bool {

	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// 判断 c 是否是字符或者数字
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}