/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
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
    public ListNode mergeKLists(ListNode[] lists) {
        if(lists == null || lists.length == 0 ){
            return null;
        }
        ListNode unsortedList = new ListNode(0);
        ListNode tmpunsorted = unsortedList;
        for (int i = 0; i< lists.length; i++) {
            if (lists[i] != null) {
                ListNode tmp = lists[i];
                while(tmp != null) {
                    tmpunsorted.next = tmp;
                    tmpunsorted = tmpunsorted.next;
                    tmp = tmp.next;
                }
            }
        }
        boolean swapped = true;
        ListNode sortedList = unsortedList;
        while(swapped) {
            unsortedList = sortedList.next;
            if (unsortedList == null) {
                return null;
            }
            swapped = false;
            while(unsortedList.next != null) {
                if (unsortedList.next.val < unsortedList.val) {
                    ListNode tmp = new ListNode(unsortedList.val);
                    tmp.next = unsortedList.next.next;
                    unsortedList.val =  unsortedList.next.val;
                    unsortedList.next = tmp;
                    swapped = true;
                }
                unsortedList = unsortedList.next;
            }
        }
        return sortedList.next;
    }
}