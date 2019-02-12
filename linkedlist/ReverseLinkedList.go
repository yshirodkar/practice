/*
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var tmpList *ListNode
    for head != nil {
        tmpNode := head.Next
        head.Next = tmpList
        tmpList = head
        head = tmpNode
    }
    return tmpList
}

// Java Solution 

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        if(head==null||head.next==null)
            return head;
 
        ListNode tmpList = null;
        while(head!=null){
            ListNode t = head.next;
            head.next = tmpList;
            tmpList = head;
            head = t;
        }

        return tmpList;
    }
}

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        if(head==null||head.next==null)
            return head;
 
        ListNode p1 = head;
        ListNode p2 = p1.next;

        head.next = null;
        while(p1!=null&& p2!=null){
            ListNode t = p2.next;
            p2.next = p1;
            p1 = p2;
            p2 = t;
        }

        return p1;
    }
}
