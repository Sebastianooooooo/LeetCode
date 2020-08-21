package main

func main() {
	
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
func threeSum(height []int) int {
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
