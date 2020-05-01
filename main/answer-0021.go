package main

import "fmt"

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 21. 合并两个有序链表
//
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
//	示例：
//	输入：1->2->4, 1->3->4
//	输出：1->1->2->3->4->4
func main() {
	listNode1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	listNode2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(&listNode1, &listNode2)
	traverseLinkedListByIteration(result)
}

// 迭代方式，遍历链表
func traverseLinkedListByIteration(result *ListNode) {
	for d := result; d != nil; d = d.Next {
		fmt.Print(d.Val)
	}
}

// 递归方式，遍历链表
func traverseLinkedListByRecursion(result *ListNode) {
	if result == nil {
		return
	}
	fmt.Print(result.Val)
	traverseLinkedListByRecursion(result.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var t ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		t.Val = l1.Val
		if l1.Next != nil {
			t.Next = mergeTwoLists(l1.Next, l2)
		} else {
			t.Next = l2
		}
	} else {
		t.Val = l2.Val
		if l2.Next != nil {
			t.Next = mergeTwoLists(l1, l2.Next)
		} else {
			t.Next = l1
		}
	}
	return &t
}
