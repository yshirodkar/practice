/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.
*/

class Solution {
    public boolean isIsomorphic(String s, String t) {
        Map<String, Integer> firstStringmap = new HashMap<>();
        Map<String, Integer> secondStringmap = new HashMap<>();
        boolean res = false;
        char[] schar = s.toCharArray();
        char[] tchar = t.toCharArray();
        if (schar.length == 0 && tchar.length == 0) {
            return true;
        }
        for (int i=0; i < schar.length; i++) {
            String tmps = String.valueOf(schar[i]);
            String tmpt = String.valueOf(tchar[i]);
            if (firstStringmap.containsKey(tmps) || secondStringmap.containsKey(tmpt) ) {
                if (firstStringmap.get(tmps) == secondStringmap.get(tmpt)) {
                    res = true;
                    continue;
                } else {
                    return false;
                }
            } else {
                res = true;
                firstStringmap.put(tmps, i);
                secondStringmap.put(tmpt, i);
            }
        }
        return res;
    }
}