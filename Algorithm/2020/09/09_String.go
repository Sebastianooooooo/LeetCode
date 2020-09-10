package main

import "fmt"

func main() {
	reverseString([]byte("H,a,n,n,a,h"))
}

/**
* 问题
* 题目要求我们反转一个字符串。
* 2020-09-09
* Example:
* Input: ["H","a","n","n","a","h"]
* Output: ["h","a","n","n","a","H"]
* OR
* Input: ["h","e","l","l","o"]
*Output: ["o","l","l","e","h"]
*
 */
func reverseString(s []byte) {

	//
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println(string(s))
}
