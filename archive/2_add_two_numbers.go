package archive

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解题思路：最低位对齐、每个位置相加进位
// 为了运算方便，后边高位直接
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result *ListNode // 结果保存
		parent *ListNode // 每轮计算用来挂载的父节点
		child  *ListNode // 每轮计算产生的新节点
		sum    int       // 当前位求和
		carry  int       // 进位值
	)
	for {
		// 死循环首先确定退出条件
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		// 计算低位,标记进位
		child = &ListNode{}
		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + carry
		} else if l1 != nil && l2 == nil {
			sum = l1.Val + carry
		} else if l1 == nil && l2 != nil {
			sum = l2.Val + carry
		} else {
			sum = carry
		}
		carry = sum / int(10)
		child.Val = sum % 10
		// 每轮计算完毕应当向后移动
		if parent == nil {
			parent = child
			result = parent
		} else {
			parent.Next = child
			parent = child
		}
		// 还能继续计算则抛弃低位
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return result
}
