/*
Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"
*/

/*
Your return is string so for fast opertaion use stringbuilder
Define lenghts and carry
run while loop till carry becomes 0 and all the digits are considered
if digit is not there then we are using 0 as a placeholder otherwise do char subtraction which will result in int
*/
public class Solution {
    public String addBinary(String a, String b) {
        StringBuilder sum = new StringBuilder();
        int i = a.length() - 1;
        int j = b.length() - 1;
        int carry = 0;
        while (i >= 0 || j >= 0 || carry == 1) {
            int digitA = i < 0 ? 0 : a.charAt(i--) - '0';
            int digitB = j < 0 ? 0 : b.charAt(j--) - '0';
            sum.insert(0, (digitA + digitB + carry) % 2);
            carry = (digitA + digitB + carry) / 2;
        }
        return sum.toString();
    }
}