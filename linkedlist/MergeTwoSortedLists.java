/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if(l1 == null) return l2;
        if(l2 == null) return l1;

        ListNode t1 = l1;
        ListNode t2 = l2;
        
        if(t1.val > t2.val) {
            ListNode tmpNode = t1;
            t1 = t2;
            t2 = tmpNode;
        }
        l1 = t1;
        
        while(t1.next != null && t2 != null) {
            if (t1.val <= t2.val && t1.next.val <= t2.val) {
                t1 = t1.next;
            } else if (t1.val <= t2.val) {
                ListNode tmpNode = t2;
                t2 = t1.next;
                t1.next = tmpNode;
            }
        }
        
        if(t1.next == null) {
            t1.next = t2;
        }
        return l1;
    }
}