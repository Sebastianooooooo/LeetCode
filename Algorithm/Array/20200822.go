package main

import (
	"fmt"
	"sort"
)

func main() {
	r1 := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(r1)
}



/**
 * 问题1
 * 给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
 * 2020-08-22
 * Example:
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * A solution set is:
	[
		[-1, 0, 1],
		[-1, -1, 2]
   ]
 *
 *
*/
func threeSum(nums []int) [][]int {

	var res [][]int
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}
	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 { //有3个值都是0的
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res

}



/**
 * 问题2
 * 给定一个数组，要求在这个数组中找出 3 个数之和离 target 最近。
 * 2020-08-22
 * Example:
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */
func threeSumClosest(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}




/**
 * 问题3
 * 给定一个数组，要求在这个数组中找出 4 个数之和为 0 的所有组合。
 * 2020-08-22
 * Example:
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 * A solution set is:
	[
  		[-1,  0, 0, 1],
  		[-2, -1, 1, 2],
  		[-2,  0, 0, 2]
	]
 *
 *
*/
func fourSum(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
