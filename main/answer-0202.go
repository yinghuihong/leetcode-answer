package main

import "fmt"

// https://leetcode-cn.com/problems/happy-number/
// 202. 快乐数
//
// 编写一个算法来判断一个数 n 是不是快乐数。
// 「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
// 如果 n 是快乐数就返回 True ；不是，则返回 False 。
//	
//	示例：
//	输入：19
//	输出：true
//	解释：
//	12 + 92 = 82
//	82 + 22 = 68
//	62 + 82 = 100
//	12 + 02 + 02 = 1
func main() {
	fmt.Println(isHappy(13))
}

// 快乐数
func isHappy(n int) bool {
	// leetcode 不让定义全局变量
	var history []int
	t := n
	for ; ; {
		t = bitSquareSum(t)
		if t == 1 {
			return true
		}
		// 遍历历史记录
		for _, h := range history {
			if h == t {
				return false
			}
		}
		// 添加记录
		history = append(history, t)
	}
}

// 拆分，计算平方和
func bitSquareSum(n int) int {
	sumOfSquares := 0
	for nextNum := n; nextNum != 0; nextNum /= 10 {
		mod := nextNum % 10
		sumOfSquares += mod * mod
		fmt.Println(mod)
	}
	fmt.Println(sumOfSquares)
	return sumOfSquares
}
