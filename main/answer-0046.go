package main

import "fmt"

// https://leetcode-cn.com/problems/permutations/
// 46. 全排列
//
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var store [][]int
	pathNums := make([]int, len(nums))
	trace(nums, pathNums, &store)
	return store
}

func trace(canChoiceNums []int, pathNums []int, store *[][]int) {
	if len(canChoiceNums) == 0 {
		pathNumsCopy := make([]int, len(pathNums))
		copy(pathNumsCopy, pathNums)
		*store = append(*store, pathNumsCopy)
		// fmt.Print("store trace ")
		// fmt.Println(store)
		return
	}
	for i, num := range canChoiceNums {
		// 将记录往后移动1格
		copy(pathNums[1:], pathNums[0:len(pathNums)-1])
		// 记录路径
		pathNums[0] = num

		// fmt.Print("=== ")
		// fmt.Print(num)
		// fmt.Println(pathNums)

		// 将num从可选数组中移除
		tmp := make([]int, len(canChoiceNums)-1)
		copy(tmp[0:], canChoiceNums[0:i])
		copy(tmp[i:], canChoiceNums[i+1:])

		// 递归
		trace(tmp, pathNums, store)

		// 将记录往前移动1格，恢复
		copy(pathNums[0:], pathNums[1:])
	}
}
