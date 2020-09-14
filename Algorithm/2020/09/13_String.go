package main

import "fmt"

func main() {
	r1 := isLongPressedName("alex", "aaleex")
	fmt.Println(r1)
}

/**
* 问题
* 给定 2 个字符串，后者的字符串中包含前者的字符串。
* 比如在打字的过程中，某个字符会多按了几下。判断后者字符串是不是比前者字符串存在这样的“长按”键盘的情况。
* 2020-09-13
* Example:
* Input: name = "alex", typed = "aaleex"
* Output: true
* Explanation: 'a' and 'e' in 'alex' were long pressed.
* Input: name = "leelee", typed = "lleeelee"
* Output: true
* Input: name = "laiden", typed = "laiden"
* Output: true
* Explanation: It's not necessary to long press any character.
*
*
 */
func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 && len(typed) == 0 {
		return true
	}
	if (len(name) == 0 && len(typed) != 0) || (len(name) != 0 && len(typed) == 0) {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		for i < len(name) && j < len(typed) && name[i] == typed[j] {
			i++
			j++
		}
		for j < len(typed) && typed[j] == typed[j-1] {
			j++
		}
	}
	return i == len(name) && j == len(typed)
}
