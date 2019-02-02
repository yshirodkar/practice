/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /* Solution 1 
28 ms
4.3 MB
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{} 
    out := head
    carry := 0
    sum := 0
    for l1 != nil && l2 != nil {
        sum = l1.Val + l2.Val + carry
        if sum > 9 {
            sum -= 10
            carry = 1
        } else {
            carry = 0
        }
        out.Next = &ListNode{Val: sum} 
        l1, l2, out = l1.Next, l2.Next, out.Next 
    }
    
    if l1 == nil && l2 == nil {
        if carry > 0 {
            out.Next = &ListNode{Val: 1}
        }
        return head.Next
    }

    // handle a remaining list
    var remaining *ListNode
    if l1 != nil {
        remaining = l1
    } else {
        remaining = l2
    }

    for carry > 0 || remaining != nil {
        if carry == 0 {
            out.Next = remaining
            return head.Next
        }
        if remaining == nil {
            out.Next = &ListNode{Val: 1}
            return head.Next
        }
        
        sum := remaining.Val + carry
        if sum > 9 {
            sum -= 10
            carry = 1
        } else {
            carry = 0
        }
        out.Next = &ListNode{Val: sum}
        remaining, out = remaining.Next, out.Next
    }

    return head.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* Solution 2 (optimised)
16 ms
4.3 MB
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := ListNode{0, nil}
    current := &head
    
    p := l1
    q := l2
    carry := 0
    
    for {
        v1 := 0
        v2 := 0
        
        if p != nil {
            v1 = p.Val
            p = p.Next
        }
        
        if q != nil {
            v2 = q.Val
            q = q.Next
        }
        
        sum := carry + v1 + v2
        carry = sum/10
        current.Next = &ListNode{sum % 10, nil}
        current = current.Next
        if p == nil && q == nil {
            break
        }
    }
    
    if carry != 0 {
        current.Next = &ListNode{carry, nil}
    }
    
    return head.Next
}