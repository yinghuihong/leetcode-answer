package main

import "fmt"

// https://leetcode-cn.com/problems/jump-game-ii/
// 45. 跳跃游戏 II
//
// 给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//	示例:
//
//	输入: [2,3,1,1,4]
//	输出: 2
//	解释: 跳到最后一个位置的最小跳跃数是 2。
//	     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//	说明:
//
//	假设你总是可以到达数组的最后一个位置。
func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

//	方法一：反向查找出发位置
//
//	复杂度分析
//
//	时间复杂度：O(n^2)，其中 n 是数组长度。有两层嵌套循环，在最坏的情况下，例如数组中的所有元素都是 1，position 需要遍历数组中的每个位置，对于 position 的每个值都有一次循环。
//
//	空间复杂度：O(1)。
func jump(nums []int) int {
	position := len(nums) - 1
	step := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				step++
				break
			}
		}
	}
	return step
}
