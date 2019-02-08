/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    rootP := head
    current := head
    
    for i := 0; i < n; i++ {
        if current.Next == nil {
            return head.Next
        }
        current = current.Next
    }
    for current.Next != nil {
        rootP, current = rootP.Next, current.Next
    }
    rootP.Next = rootP.Next.Next
    return head
}