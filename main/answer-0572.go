package main

import "fmt"

// https://leetcode-cn.com/problems/subtree-of-another-tree/
// 572. 另一个树的子树
//
// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。
//
//	示例 1:
//	给定的树 s:
//
//	    3
//	   / \
//	  4   5
//	 / \
//	1   2
//	给定的树 t：
//
//	  4
//	 / \
//	1   2
//	返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
//
//	示例 2:
//	给定的树 s：
//
//	      3
//	     / \
//	    4   5
//	   / \
//	  1   2
//	 /
//	0
//	给定的树 t：
//
//	  4
//	 / \
//	1   2
//	返回 false。
func main() {
	//var s = TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 1,
	//			Left: &TreeNode{
	//				Val:   0,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   2,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   5,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//var t = TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 1,
	//		Left: &TreeNode{
	//			Val:   0,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	var s = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	var t = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSubtree(&s, &t))
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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}
