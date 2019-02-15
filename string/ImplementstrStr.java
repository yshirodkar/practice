/*
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:
What should we return when needle is an empty string? This is a great question to ask during an interview.
For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

/*
First check the needle length is 0 or not
Then check for haystack lenght.
If haystacka and needle length is same the check both char are same or not
Run loop from 0 to haystack length - lenght of needle +1
Use equals function to check and do substring operation on haystach.
*/
class Solution {
    public int strStr(String haystack, String needle) {
        if(needle.length() == 0) return 0;
        if(haystack.length() == 0) return -1;
        if(haystack.length() == 1 && needle.length() == 1){
            if(haystack.charAt(0) == needle.charAt(0)){
                return 0;
            }
        }
        for (int i = 0; i < haystack.length() - needle.length() + 1; i++) {
            if (needle.equals(haystack.substring(i, i + needle.length()))) {
                return i;
            }
        }
        return -1;
    }
}