package main

import (
	"fmt"
	"strings"
)

func main() {
	r1 := isPalindrome8("A man, a plan, a canal: Panama")
	fmt.Println(r1)
}


/**
* 问题
* 判断所给的字符串是否是有效的回文串。
* 2020-09-08
* Example:
* "A man, a plan, a canal: Panama" is a palindrome.
* "race a car" is not a palindrome.
*
*
*
*/
func isPalindrome8(s string) bool {

	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar8(s[i]) {
			i++
		}
		for i < j && !isChar8(s[j]) {
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
func isChar8(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
