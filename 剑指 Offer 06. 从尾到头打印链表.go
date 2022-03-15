/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++ //记录长度
	}
	var res []int = make([]int, l)
	for head != nil {
		res[l-1] = head.Val //从尾到头存放val
		head = head.Next
		l--
	}
	return res

}