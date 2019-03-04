class Solution {
	public boolean isPalindrome(ListNode head) {
		if (head == null) {
			return true;
		}

		ListNode p = head;
		ListNode prev = new ListNode(head.val);
		while(p.next != null) {
			ListNode tmp = new ListNode (p.next.val);
			tmp.next = prev;
			prev = tmp;
			p = p.next;
		}
	}
}