package main

import (
	"fmt"
	"sort"
)

func main() {
	r := reorganizeString("aab")
	fmt.Println(r)
}

/**
* 问题
* 给定一个字符串，要求重新排列字符串，让字符串两两字符不相同，如果可以实现，即输出最终的字符串，如果不能让两两不相同，则输出空字符串。
* 2020-09-11
* Example:
* Input: S = "aab"
* Output: "aba"
* OR
* Input: S = "aaab"
* Output: ""
*
 */
func reorganizeString(S string) string {
	fs := frequencySort767(S)
	if fs == "" {
		return ""
	}
	bs := []byte(fs)
	ans := ""
	j := (len(bs)-1)/2 + 1
	for i := 0; i <= (len(bs)-1)/2; i++ {
		ans += string(bs[i])
		if j < len(bs) {
			ans += string(bs[j])
		}
		j++
	}
	return ans
}

func frequencySort767(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
		if sMap[b] > (len(sb)+1)/2 {
			return ""
		}
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}
