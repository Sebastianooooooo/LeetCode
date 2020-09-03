package main

import (
	"fmt"
)

func main() {
	r1 := removeDuplicates([]int{1,1,2})
	fmt.Println(r1)
	r2 := removeElement([]int{0,1,2,2,3,0,4,2},2)
	fmt.Println(r2)
}


/**
 * 问题1
 * 给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值。
 * 2020-08-23
 * Example:
 * Given nums = [1,1,2],
 * Your function should return length = 2,
 * Given nums = [0,0,1,1,1,2,2,3,3,4],
 * Your function should return length = 5
 *
 *
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1

}



/**
 * 问题2
 * 给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。
 * 2020-08-23
 * Example:
 * Given nums = [0,1,2,2,3,0,4,2], val = 2,
 * Your function should return length = 5,
 * 这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组实际剩余的元素个数，
 * OJ 最终判断题目的时候会读取数组剩余个数的元素进行输出。
 *
 */
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	fmt.Println(nums)
	return j
}