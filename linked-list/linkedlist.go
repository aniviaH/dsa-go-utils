package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

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
