/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*/



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
 8 ms
 3.7MB
 */
func removeElements(head *ListNode, val int) *ListNode {
    returnList := &ListNode{}
    returnList.Next = head
    preList := returnList
    for preList.Next != nil {
        if preList.Next.Val == val {
            preList.Next = preList.Next.Next
        } else {
            preList = preList.Next
        }
    }
    return returnList.Next
}