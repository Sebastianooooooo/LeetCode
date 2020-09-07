package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	r1 := simplifyPath("/a/../../b/../c//.//")
	fmt.Println(r1)

}

/**
* 问题
* 给出一个 Unix 的文件路径，要求简化这个路径。这道题也是考察栈的题目。
* 2020-09-05
* Example:
* Input: "/a/../../b/../c//.//"
* Output: "/c"
*
* Input: "/a//b////c/d//././/.."
* Output: "/a/b/c"
*
 */

// 解法一
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	var res string
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		//cur := strings.TrimSpace(arr[i]) 更加严谨的做法应该还要去掉末尾的空格
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}

// 解法二 golang 的官方库 API
func simplifyPath1(path string) string {
	return filepath.Clean(path)
}
