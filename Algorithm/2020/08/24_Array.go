package main

import (
	"fmt"
	"sort"
)

func main() {
	r1 := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(r1)
	r2 := combinationSum2([]int{2,5,2,1,2}, 5)
	fmt.Println(r2)
}

/**
 * 问题1
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * candidates 中的数字可以无限制重复被选取。
 * 2020-08-24
 * Example:
 * Input: candidates = [2,3,6,7], target = 7,
	A solution set is:
	[
		[7],
		[2,2,3]
	]
 * 题目要求出总和为 sum 的所有组合，组合需要去重。
 * Your function should return length = 5
 *
 *
*/
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target { // 这里可以剪枝优化
			break
		}
		c = append(c, nums[i])
		findCombinationSum(nums, target-nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}



/**
 * 问题2
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * candidates 中的每个数字在每个组合中只能使用一次。
 * 2020-08-24
 * Example:
 * Input: candidates = [2,5,2,1,2], target = 5,
	A solution set is:
	[
		[1,2,2],
  		[5]
	]
 * 题目要求出总和为 sum 的所有组合，组合需要去重。
 * Your function should return length = 5
 *
 *
*/
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates) // 这里是去重的关键逻辑
	findCombinationSum2(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
			continue
		}
		if target >= nums[i] {
			c = append(c, nums[i])
			findCombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
