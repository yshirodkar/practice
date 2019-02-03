/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

/*
8 ms
3.3 MB
*/
func Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func lengthOfLongestSubstring(s string) int {
    m := make(map[rune]int)
    i := 0
    maxLen := 0
    ok := false // this asved 4 ms for total run
    ss := []rune(s)
    for j := 0; j < len(ss); j++ {
        _, ok = m[ss[j]]
        if ok { 
           i =  Max(m[ss[j]], i)
        }
        maxLen = Max(maxLen, j-i+1)
        m[ss[j]] = j + 1
    }
    return maxLen
}