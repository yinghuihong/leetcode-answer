package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 98. 验证二叉搜索树
//
//	给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//	假设一个二叉搜索树具有如下特征：
//
//	节点的左子树只包含小于当前节点的数。
//	节点的右子树只包含大于当前节点的数。
//	所有左子树和右子树自身必须也是二叉搜索树。
//	示例 1:
//
//	输入:
//	  2
//	 / \
//	1   3
//	输出: true
//	示例 2:
//
//	输入:
//	  5
//	 / \
//	1   4
//	   / \
//	  3   6
//	输出: false
//	解释: 输入为: [5,1,4,null,null,3,6]。
//	     根节点的值为 5 ，但是其右子节点值为 4 。
func main() {
	var root = TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isValidBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper(root, -math.MaxInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	fmt.Print(lower)
	fmt.Print(" ")
	fmt.Println(upper)
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
