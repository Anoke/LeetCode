package main

import "LeetCode/Easy"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummy.Next
}
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := make([]int, 128)
	map2 := make([]int, 128)

	for i := 0; i < len(s); i++ {
		sch := s[i]
		tch := t[i]

		if map1[sch] == 0 && map2[tch] == 0 {
			map1[sch] = int(tch)
			map2[tch] = int(sch)
		} else if map1[sch] != int(tch) || map2[tch] != int(sch) {
			return false
		}
	}
	return true
}
func main() {
	Easy.RomanToInt("IV")
}
