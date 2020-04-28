package main

import "fmt"

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 23. 合并K个排序链表
//
// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
// 示例:
// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6
func main() {
	listNode1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
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
	listNode3 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	lists := []*ListNode{
		&listNode1, &listNode2, &listNode3,
	}
	result := mergeKLists(lists)
	traverseLinkedListByIteration(result)
	fmt.Println()
	traverseLinkedListByRecursion(result)
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
func mergeKLists(lists []*ListNode) *ListNode {
	var tmp *ListNode
	for _, listNode := range lists {
		if listNode == nil {
			continue
		}
		if tmp == nil {
			tmp = listNode
			continue
		}
		tmp = merge(tmp, listNode)
	}
	//traverseLinkedListByRecursion(tmp)
	return tmp
}

func merge(a *ListNode, b *ListNode) *ListNode {
	//traverseLinkedListByRecursion(a)
	//traverseLinkedListByRecursion(b)
	var t ListNode
	if a.Val < b.Val {
		t.Val = a.Val
		if a.Next != nil {
			t.Next = merge(a.Next, b)
		} else {
			t.Next = b
		}
	} else {
		t.Val = b.Val
		if b.Next != nil {
			t.Next = merge(a, b.Next)
		} else {
			t.Next = a
		}
	}
	return &t
}
