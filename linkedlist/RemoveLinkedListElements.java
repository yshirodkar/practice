/*
Remove all elements from a linked list of integers that have value val.

Example:
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5

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
    public ListNode removeElements(ListNode head, int val) {
        ListNode helper = new ListNode(0);
        helper.next = head;
        ListNode tmp = helper;
        while ( tmp.next != null) {
            if (tmp.next.val == val) {
                if (tmp.next.next != null) {
                    ListNode next = tmp.next;
                    tmp.next = next.next;
                } else {
                    tmp.next = null;
                }
                
            } else {
                tmp = tmp.next;
            }
        }
        
        return helper.next;
    }
}