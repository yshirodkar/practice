/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

/*
56 ms
5.7 MB
*/
func isPalindrome(x int) bool {
    if  x < 0 {
        return false
    }
    var res []int
    i := 0
    if x == 0 {
        return true
    }
    for x != 0 {
        res = append(res,  x%10)
        x = x/10
        i++
    }
    if len(res) < 3 {
        if (len(res) == 2 && (res[0] == res[1])) || len(res) == 1{
            return true
        }
        return false
    }
    
    for i:= 0; i < len(res)/2; i++ {
        if res[i] != res[len(res) - i - 1] {
            return false
        }
    }
    return true
}