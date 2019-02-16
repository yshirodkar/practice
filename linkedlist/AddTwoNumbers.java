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
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int carry = 0;
        int sum = 0;
        ListNode returnval = new ListNode(0);
        ListNode res = returnval;
        while(l1 != null && l2 != null) {
            sum = l1.val + l2.val + carry;
            if (sum > 9) {
                sum = sum - 10;
                carry = 1;
            } else {
                carry = 0;
            }
            res.next = new ListNode(sum);
            l1 = l1.next;
            l2 = l2.next;
            res = res.next;
        }

        if (l1 == null && l2 == null) {
            if (carry ==1) {
                res.next = new ListNode(carry);
            }
            return returnval.next;
        }
        ListNode remain = new ListNode(0);
        if (l1 != null) {
            remain = l1;
        } else {
            remain = l2;
        }
        
        while(carry > 0 || remain != null) {
            if (carry == 0) {
                res.next = remain;
                return returnval.next;
            }
            
            if (remain == null) {
                res.next = new ListNode(carry);
                return returnval.next;
            }
            sum = remain.val + carry;
            if (sum > 9) {
                sum = sum - 10;
                carry = 1;
            } else {
                carry = 0;
            }
            res.next = new ListNode(sum);
            remain = remain.next;
            res = res.next;
        }
        return returnval.next;
    }
}