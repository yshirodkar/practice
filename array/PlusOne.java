/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:
Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.

Example 2:
Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.

*/
class Solution {
    public int[] plusOne(int[] digits) {
        if(digits==null||digits.length==0)
            return new int[0];
        
        int carry = 1;
        for (int i=digits.length-1; i >=0; i--) {
            int tmp = digits[i] + carry;
            if (tmp > 9) {
                digits[i] = 1;
                carry = 1;
            } else {
                carry = 0;
            }
            digits[i] = tmp%10;
        }
        
        if (carry == 1) {
            int[] newresult = new int[digits.length + 1];
            System.arraycopy(digits, 0, newresult, 1, digits.length);
            newresult[0] = 1;
            return newresult;
        } else {
            return digits;
        }
    }
}