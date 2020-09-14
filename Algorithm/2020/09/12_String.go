package main

import "fmt"

func main() {
	r1 := scoreOfParentheses("()()")
	fmt.Println(r1)
}


/**
* 问题
* 按照以下规则计算括号的分数：() 代表 1 分。AB 代表 A + B，A 和 B 分别是已经满足匹配规则的括号组。
* (A) 代表 2 * A，其中 A 也是已经满足匹配规则的括号组。给出一个括号字符串，
* 要求按照这些规则计算出括号的分数值。
* 2020-09-12
* Example:
* Input: "()()"
* Output: 2
* Input: "(()(()))"
* Output: 6
*
*
*/
func scoreOfParentheses(S string) int {
	res, stack, top, temp := 0, []int{}, -1, 0
	for _, s := range S {
		if s == '(' {
			stack = append(stack, -1)
			top++
		} else {
			temp = 0
			for stack[top] != -1 {
				temp += stack[top]
				stack = stack[:len(stack)-1]
				top--
			}
			stack = stack[:len(stack)-1]
			top--
			if temp == 0 {
				stack = append(stack, 1)
				top++
			} else {
				stack = append(stack, temp*2)
				top++
			}
		}
	}
	for len(stack) != 0 {
		res += stack[top]
		stack = stack[:len(stack)-1]
		top--
	}
	return res
}

