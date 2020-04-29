package main

import "fmt"

// https://leetcode-cn.com/problems/find-in-mountain-array/
// 1095. 山脉数组中查找目标值
//
// 给你一个 山脉数组 mountainArr，请你返回能够使得 mountainArr.get(index) 等于 target 最小 的下标 index 值。
// 如果不存在这样的下标 index，就请返回 -1。
//
//
// 何为山脉数组？如果数组 A 是一个山脉数组的话，那它满足如下条件：
// 首先，A.length >= 3
// 其次，在 0 < i < A.length - 1 条件下，存在 i 使得：
// A[0] < A[1] < ... A[i-1] < A[i]
// A[i] > A[i+1] > ... > A[A.length - 1]
// 
//
// 你将 不能直接访问该山脉数组，必须通过 MountainArray 接口来获取数据：
//
// MountainArray.get(k) - 会返回数组中索引为k 的元素（下标从 0 开始）
// MountainArray.length() - 会返回该数组的长度
//
//	示例 1：
//	输入：array = [1,2,3,4,5,3,1], target = 3
//	输出：2
//	解释：3 在数组中出现了两次，下标分别为 2 和 5，我们返回最小的下标 2。
//
//	示例 2：
//	输入：array = [0,1,2,4,2,1], target = 3
//	输出：-1
//	解释：3 在数组中没有出现，返回 -1。
func main() {
	ma := MountainArray{array: []int{3, 5, 3, 2, 0}}
	target := 2
	fmt.Println(findInMountainArray(target, &ma))
}

type MountainArray struct {
	array []int
}

func (ma *MountainArray) get(index int) int {
	return ma.array[index]
}

func (ma *MountainArray) length() int {
	return len(ma.array)
}

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, mountainArr *MountainArray) int {
	// 找到最高点
	var mountainIndex int
	total := mountainArr.length()
	left, right := 1, total-2
	for left <= right {
		mid := (left + right) / 2
		//fmt.Println(mid)
		if left == right {
			mountainIndex = left
			break
		} else if mountainArr.get(mid-1) < mountainArr.get(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 左上坡
	left = 0
	right = mountainIndex
	for left <= right {
		mid := (left + right) / 2
		if mountainArr.get(mid) == target {
			return mid
		} else if mountainArr.get(mid) < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 右下坡
	left = mountainIndex
	right = total - 1
	for left <= right {
		mid := (left + right) / 2
		if mountainArr.get(mid) == target {
			return mid
		} else if mountainArr.get(mid) > target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
