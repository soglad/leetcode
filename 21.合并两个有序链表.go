/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (54.44%)
 * Likes:    468
 * Dislikes: 0
 * Total Accepted:    71.1K
 * Total Submissions: 130.4K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 
 * 示例：
 * 
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			if cur == nil { //head not set yet
				head = l1
				cur = l1
			}else {
				cur.Next = l1
				cur = cur.Next
			}
			l1 = l1.Next
		}else {
			if cur == nil { //head not set yet
				head = l2
				cur = l2
			}else {
				cur.Next = l2
				cur = cur.Next
			}
			l2 = l2.Next
		}
	}
	var t *ListNode
	if l1 != nil {
		t=l1
	}else {
		t=l2
	}
	if cur == nil {
		head = t
	}else{
		cur.Next=t
	}
	return head
}

