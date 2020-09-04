package main

import (
	"fmt"
	"sort"
)

func main() {
	r1 := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(r1)
}

/**
* 问题1
* 给出一个字符串数组，要求对字符串数组里面有 Anagrams 关系的字符串进行分组。
* Anagrams 关系是指两个字符串的字符完全相同，顺序不同，两者是由排列组合组成。
* 2020-09-04
* Example:
* Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
* Output:
 [
	  ["ate","eat","tea"],
	  ["nat","tan"],
	  ["bat"]
 ]
*
*
*/
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
