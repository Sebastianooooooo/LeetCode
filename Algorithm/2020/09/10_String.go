package main

import "fmt"

func main() {
	r := reverseVowels("leetcode")
	fmt.Println(r)
}

/**
* 问题
* 题目要求我们反转字符串中的元音字母。需要注意字母大小写。
* 2020-09-10
* Example:
* Input: "leetcode"
* Output: "leotcede"
* OR
* Input: "hello"
* Output: "holle"
*
 */
func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if isVowels(b[i]) && isVowels(b[j]) {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		} else if isVowels(b[i]) && !isVowels(b[j]) {
			j--
		} else if !isVowels(b[i]) && isVowels(b[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	return string(b)
}

func isVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
