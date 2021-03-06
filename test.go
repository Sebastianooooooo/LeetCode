package main

import "fmt"

var nil = new(int)

func main() {
	var p *int
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

//
//func main() {
//	r1 := removeDuplicates([]int{1,1,2})
//	fmt.Println(r1)
//	r2 := removeElement([]int{0,1,2,2,3,0,4,2},2)
//	fmt.Println(r2)
//}
//
//
///**
// * 问题
// * 给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值。
// * 2020-09-07
// * Example:
// * Given nums = [1,1,2],
// * Your function should return length = 2,
// * Given nums = [0,0,1,1,1,2,2,3,3,4],
// * Your function should return length = 5
// *
// *
// */
//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	last, finder := 0, 0
//	for last < len(nums)-1 {
//		for nums[finder] == nums[last] {
//			finder++
//			if finder == len(nums) {
//				return last + 1
//			}
//		}
//		nums[last+1] = nums[finder]
//		last++
//	}
//	return last + 1
//
//}
//
