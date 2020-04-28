package main

import "fmt"

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
// 搜索旋转排序数组
//
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
// 你可以假设数组中不存在重复的元素。
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 示例 1:
// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
//
// 示例 2:
// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	total := len(nums)
	left := 0
	right := total - 1
	for ; left <= right; {
		mid := (left + right) / 2
		// fmt.Println(mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			// 右半边有序
			if nums[mid] < target && target <= nums[right] {
				// 并且目标值在有序的右半边
				left = mid + 1
			} else {
				// 目标值不在有序的右半边，将范围指向无序的左半边
				right = mid - 1
			}
		} else {
			// 左半边有序
			if nums[left] <= target && target < nums[mid] {
				// 并且目标值在有序的左半边
				right = mid - 1
			} else {
				// 目标值不在有序的左半边，将范围指向无序的右半边
				left = mid + 1
			}
		}
	}
	return -1
}
