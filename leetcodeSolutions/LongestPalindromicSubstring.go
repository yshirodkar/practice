/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/


/*
76 ms
1.2 MB
*/
func ispalindrome(sub_string string) bool {
    for ind := 0; ind < (len(sub_string)/2); ind++ {
        if (sub_string[ind] != sub_string[len(sub_string) - ind - 1]) {
            return false
        }
    }
    return true
}


func longestPalindrome(s string) string {
    sub_string := ""
	if len(s) < 3 {
        if ispalindrome(s) {
    		return s
        } else if len(s) == 2 || len(s) == 1 {
            return string(s[0])
        } else {
            return s
        }
	}

    sub_string = string(s[0])
    for ind1 := 0; ind1 < len(s); ind1++ {
        subStringLen := len(sub_string)
        for ind2 := ind1+subStringLen; ind2 < len(s); ind2 ++ {
            if s[ind1] == s[ind2] {
                if ispalindrome(string(s[ind1: ind2+1])) { // array takes values from inside the range e.g. s[0:2] means s[0], s[1]
                    if len(string(s[ind1: ind2+1])) >= len(sub_string) {
                        sub_string = string(s[ind1: ind2+1])
                    }
                }
            }
        }
    }
    
    return sub_string
}