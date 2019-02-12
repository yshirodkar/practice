/*
Given a singly linked list, determine if it is a palindrome.

Example 1:
Input: 1->2
Output: false

Example 2:
Input: 1->2->2->1
Output: true
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
    public boolean isPalindrome(ListNode head) {
        if(head==null)
            return true;

        ListNode p = head;
        ListNode prev = new ListNode(head.val);
        while(p.next != null){
            ListNode t = new ListNode(p.next.val);
            t.next = prev;
            prev = t;
            p = p.next;
        }

        ListNode p1 = head;
        ListNode p2 = prev;
        while(p1 != null) {
            if (p1.val != p2.val) {
                return false;
            }
            p1 = p1.next;
            p2 = p2.next;
        }
        return true;
    }
}