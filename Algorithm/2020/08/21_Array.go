package main

import "fmt"

func main() {
	r1 := twoSum([]int{2, 7, 11, 15}, 13)
	fmt.Println(r1)
	r2 := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(r2)
}


/**
 * 问题1
 * 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
 * 2020-08-21
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1]
 *
 * 这道题最优的做法时间复杂度是 O(n)。
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

/**
 * 问题2
 * 给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。
 * 2020-08-21
 * Example:
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 *
 * 这一题也是对撞指针的思路。首尾分别 2 个指针，每次移动以后都分别判断长宽的乘积是否最大。
 */
func maxArea(height []int) int {
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


