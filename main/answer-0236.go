package main

import "fmt"

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
//
//	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//	百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//	例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
//
//	示例 1:
//
//	输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//	输出: 3
//	解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//	示例 2:
//
//	输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//	输出: 5
//	解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//	 
//
//	说明:
//
//	所有节点的值都是唯一的。
//	p、q 为不同节点且均存在于给定的二叉树中。
func main() {
	var rootTree = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	var p = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	var q = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(lowestCommonAncestor(rootTree, p, q).Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 *
 * 复杂度分析
 * 时间复杂度：O(N)O(N)，其中 NN 是二叉树的节点数。二叉树的所有节点有且只会被访问一次，因此时间复杂度为 O(N)O(N)。
 * 空间复杂度：O(N)O(N) ，其中 NN 是二叉树的节点数。递归调用的栈深度取决于二叉树的高度，二叉树最坏情况下为一条链，此时高度为 NN，因此空间复杂度为 O(N)O(N)。
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	fmt.Print("in ")
	fmt.Println(root.Val)
	if root.Val == p.Val || root.Val == q.Val {
		fmt.Print("---> ")
		fmt.Println(root.Val)
		return root
	}
	// before
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// after
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		fmt.Print("out ")
		fmt.Println(root.Val)
		return right
	}
	return left
}
