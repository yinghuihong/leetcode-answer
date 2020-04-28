package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
// 面试题56 - I. 数组中数字出现的次数
//
// 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
// 示例 1：
// 输入：nums = [4,1,4,6]
// 输出：[1,6] 或 [6,1]
func main() {
	nums := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(nums))
}

// 二进制位运算，用到了与（&）、异或异或（^）
func singleNumbers(nums []int) []int {
	var ret int
	for _, num := range nums {
		ret ^= num
		//fmt.Println(ret)
	}

	// 如果我们可以把所有数字分成两组，使得：
	// 1.两个只出现一次的数字在不同的组中；
	// 2.相同的数字会被分到相同的组中。
	// 那么对两个组分别进行异或操作，即可得到答案的两个数字。

	// 在异或结果中找到任意为 1 的位。
	// 根据这一位对所有的数字进行分组。
	div := 1
	//fmt.Println(ret & div)
	for ; ret&div == 0; {
		div = div << 1
	}
	//fmt.Println(div)

	var a, b int
	for _, num := range nums {
		if num&div == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

//func singleNumbers(nums []int) []int {
//	var tempNums []int
//	for _, num := range nums {
//		fmt.Println(num)
//		var exist bool
//		for i, tempNum := range tempNums {
//			if tempNum == num {
//				tempNums = append(tempNums[:i], tempNums[i+1:]...)
//				exist = true
//				break
//			}
//		}
//		if !exist {
//			tempNums = append(tempNums, num)
//		}
//	}
//	return tempNums
//}
