package main

import (
	"fmt"
	"strconv"
)

func main() {
	r1 := restoreIPAddresses("25525511135")
	fmt.Println(r1)
}

/**
* 问题
* 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
*
* Input: "25525511135"
* Output: ["255.255.11.135", "255.255.111.35"]
*
 */
func restoreIPAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	res, ip := []string{}, []int{}
	dfs(s, 0, ip, &res)
	return res
}

func dfs(s string, index int, ip []int, res *[]string) {
	if index == len(s) {
		if len(ip) == 4 {
			*res = append(*res, getString(ip))
		}
		return
	}
	if index == 0 {
		num, _ := strconv.Atoi(string(s[0]))
		ip = append(ip, num)
		dfs(s, index+1, ip, res)
	} else {
		num, _ := strconv.Atoi(string(s[index]))
		next := ip[len(ip)-1]*10 + num
		if next <= 255 && ip[len(ip)-1] != 0 {
			ip[len(ip)-1] = next
			dfs(s, index+1, ip, res)
			ip[len(ip)-1] /= 10
		}
		if len(ip) < 4 {
			ip = append(ip, num)
			dfs(s, index+1, ip, res)
			ip = ip[:len(ip)-1]
		}
	}
}

func getString(ip []int) string {
	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}
