package main

import "fmt"

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// 面试题51. 数组中的逆序对
//
// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
//
// 示例 1:
// 输入: [7,5,6,4]
// 输出: 5
func main() {
	nums := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	count := 0
	tmp := make([]int, n)
	mergeSort(nums, 0, n-1, &count, tmp)
	return count
}

//mergeSort
func mergeSort(nums []int, left, right int, count *int, tmp []int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(nums, left, mid, count, tmp)
	mergeSort(nums, mid+1, right, count, tmp)
	merge(nums, left, mid, right, count, tmp)
}

//merge
func merge(nums []int, left, mid, right int, count *int, tmp []int) {
	l, r := left, mid+1
	index := left
	for l <= mid && r <= right {
		if nums[l] <= nums[r] {
			tmp[index] = nums[l]
			//当左边需要弹出去时，更新逆序对数量
			*count += r - 1 - mid
			l++
			index++
		} else {
			tmp[index] = nums[r]
			r++
			index++
		}
	}
	for l <= mid {
		tmp[index] = nums[l]
		*count += r - 1 - mid
		l++
		index++
	}
	for r <= right {
		tmp[index] = nums[r]
		r++
		index++
	}
	for i := left; i <= right; i++ {
		nums[i] = tmp[i]
	}
}
