package main

import (
	"fmt"

	linkedlist "github.com/aniviaH/dsa-go-utils/linked-list/pkg"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4}
	slice3 := []int{7, 8, 9}

	// 使用for循环遍历切片
	// for index, value := range slice {
	// 	fmt.Println(index, value)
	// }

	// 构建链表
	listHead1 := linkedlist.NewLinkedListFromSlice(slice)
	listHead2 := linkedlist.NewLinkedListFromSlice(slice2)
	listHead3 := linkedlist.NewLinkedListFromSlice(slice3)

	// 打印链表的每一项
	fmt.Println()
	fmt.Println("======= 打印链表 =======")
	// fmt.Println(list)
	linkedlist.PrintLinkedList(listHead1)
	linkedlist.PrintLinkedList(listHead2)

	// 取中间节点方式1
	fmt.Println()
	fmt.Println("======= 获取链表的中间节点 =======")
	fmt.Println("获取链表中间节点-总数偶数时取中间两个的后者：", slice, "的中间节点为", linkedlist.GetLinkedListMidNode(listHead1))
	fmt.Println("获取链表中间节点-总数偶数时取中间两个的后者：", slice2, "的中间节点为", linkedlist.GetLinkedListMidNode(listHead2))

	// 取中间节点方式2
	fmt.Println("获取链表中间节点-总数偶数时取中间两个的前者：", slice, "的中间节点为", linkedlist.GetLinkedListMidNode2(listHead1))
	fmt.Println("获取链表中间节点-总数偶数时取中间两个的前者：", slice2, "的中间节点为", linkedlist.GetLinkedListMidNode2(listHead2))

	// 取中间节点方式3
	fmt.Println("获取链表中间节点-总数偶数时取中间两个的前者：", slice, "的中间节点为", linkedlist.GetLinkedListMidNode3(listHead1))
	fmt.Println("获取链表中间节点-总数偶数时取中间两个的前者：", slice2, "的中间节点为", linkedlist.GetLinkedListMidNode3(listHead2))

	// 链表反转
	fmt.Println()
	fmt.Println("======== 反转链表 =======")
	fmt.Println("---------链表:", slice3, "反转之前-----------")
	linkedlist.PrintLinkedList(listHead3)
	listHead3Reversed := linkedlist.ReverseLinkedList(listHead3)
	fmt.Println("---------链表:", slice3, "反转之后-----------")
	fmt.Print("原链表头：")
	linkedlist.PrintLinkedList(listHead3)
	fmt.Print("新链表头：")
	linkedlist.PrintLinkedList(listHead3Reversed)
	fmt.Println()

	// 拼接两个已排序的链表，得到一条排序的链表
	fmt.Println()
	slice4 := []int{1, 4, 7}
	slice5 := []int{3, 6, 9}
	listHead4 := linkedlist.NewLinkedListFromSlice(slice4)
	listHead5 := linkedlist.NewLinkedListFromSlice(slice5)
	fmt.Println("========合并链表", slice4, "和", slice5, "=======")
	fmt.Print("合并前的", slice4, "链表：")
	linkedlist.PrintLinkedList(listHead4)
	fmt.Print("合并后的", slice4, "链表：")
	linkedlist.PrintLinkedList(listHead5)
	mergedHead := linkedlist.MergeTwoSortedLinkedList(listHead4, listHead5)
	fmt.Print("合并后的结果链表：")
	linkedlist.PrintLinkedList(mergedHead)
	fmt.Print("合并后的", slice4, "链表：")
	linkedlist.PrintLinkedList(listHead4)
	fmt.Print("合并后的", slice5, "链表：")
	linkedlist.PrintLinkedList(listHead5)

	// 剪裁链表的前n个位置，返回剪裁之后的头部节点
	fmt.Println()
	slice6 := []int{1, 2, 3, 4, 5}
	listHead6 := linkedlist.NewLinkedListFromSlice(slice6)
	fmt.Print("裁剪前的", slice6, "链表：")
	linkedlist.PrintLinkedList(listHead6)

	resCut1 := linkedlist.CutLinkedListPreCount(listHead6, 2)
	fmt.Print("裁剪第一次后的返回结果链表：")
	linkedlist.PrintLinkedList(resCut1)
	fmt.Print("裁剪第一次后的", slice6, "链表：")
	linkedlist.PrintLinkedList(listHead6)

	resCut2 := linkedlist.CutLinkedListPreCount(resCut1, 2)
	fmt.Print("裁剪第二次后的返回结果链表：")
	linkedlist.PrintLinkedList(resCut2)
	fmt.Print("裁剪第二次后的", slice6, "链表：")
	linkedlist.PrintLinkedList(listHead6)
	fmt.Print("裁剪第二次后的resCut1链表：")
	linkedlist.PrintLinkedList(resCut1)

	resCut3 := linkedlist.CutLinkedListPreCount(resCut2, 2)
	fmt.Print("裁剪第三次后的返回结果链表：")
	linkedlist.PrintLinkedList(resCut3)
	fmt.Print("裁剪第三次后的", slice6, "链表：")
	linkedlist.PrintLinkedList(listHead6)
	fmt.Print("裁剪第三次后的resCut1链表：")
	linkedlist.PrintLinkedList(resCut1)
	fmt.Print("裁剪第三次后的resCut2链表：")
	linkedlist.PrintLinkedList(resCut2)
}
