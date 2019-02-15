/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:
All given inputs are in lowercase letters a-z.
*/

/*
If the array length is less than 1 then return "".
Set first index of array as our prefix (strs[0])
Check this prefix with next index of array if the indexof that is 0 then just increase index of array
if indexof doesn't match then reduce the prefix substring by 1 and loop again. Run this loop till prefix length > 0
return prefix.
*/
class Solution {
    public String longestCommonPrefix(String[] strs) {
        int strlenght = strs.length;
        if (strlenght < 1) {
            return "";
        }
        String prefix = strs[0].substring(0, strs[0].length());
        if (strlenght >= 2) {
            int i = 1;
            while(prefix.length() > 0) {
                if (strs[i].indexOf(prefix) == 0) {
                    i++;
                    if (strlenght == i) {
                        break;   
                    }
                } else {
                    prefix = prefix.substring(0, prefix.length() -1);
                }
            }
        }
        return prefix;
    }
}