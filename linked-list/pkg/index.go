package pkg

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 以切片构建链表
func NewLinkedListFromSlice(s []int) *ListNode {
	// 使用for循环遍历切片
	dummyNode := &ListNode{Val: -1}
	head := dummyNode
	for _, value := range s {
		node := &ListNode{Val: value}

		head.Next = node
		head = head.Next
	}

	return dummyNode.Next
}

// 打印链表
func PrintLinkedList(head *ListNode) {
	dummyNode := &ListNode{Next: head}

	cur := dummyNode.Next
	for cur != nil {
		fmt.Print(cur.Val)
		fmt.Print(" -> ")
		cur = cur.Next
	}

	fmt.Print("nil")
	fmt.Println()
}

// 获取中间节点
// 链表节点总数为奇数：中间节点
// 链表节点总数为偶数：中间的两个节点的后面节点 - 对应leetcode的876题：链表的中间节点 [https://leetcode.cn/problems/middle-of-the-linked-list/description/]
func GetLinkedListMidNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 获取中间节点
// 链表节点总数为奇数：中间节点的前面一个节点(奇数场景取到的不是预期的！)
// 链表节点总数为偶数：中间的两个节点的前面节点 - leetcode的234题：回文链表，用到奇偶场景的正确中间节点的判断 [https://leetcode.cn/problems/middle-of-the-linked-list/description/]
func GetLinkedListMidNode2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return pre
}

// 获取中间节点
// 链表节点总数为奇数：中间节点
// 链表节点总数为偶数：中间的两个节点的前面节点 -
func GetLinkedListMidNode3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 反转整个链表
func ReverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre
}

// 合并两条已排序的链表
func MergeTwoSortedLinkedList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	dummyNode := &ListNode{Val: -1}
	pre := dummyNode

	for head1 != nil && head2 != nil {
		// 谁小将pre指向谁，谁向后移动
		if head1.Val <= head2.Val {
			pre.Next = head1
			head1 = head1.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
		}

		// 挪动pre指针
		pre = pre.Next
	}

	// head1 或 head2 还剩未遍历的
	if head1 != nil {
		pre.Next = head1
	}
	if head2 != nil {
		pre.Next = head2
	}

	return dummyNode.Next
}

func CutLinkedListPreCount(head *ListNode, n int) *ListNode {
	cur := head

	cutLoopCount := n - 1
	for cutLoopCount > 0 && cur != nil {
		cur = cur.Next
		cutLoopCount--
	}

	// 如果裁剪到链表的末尾，这个时候cur为nil了，返回nil
	if cur == nil {
		return nil
	}

	next := cur.Next
	cur.Next = nil

	return next
}
