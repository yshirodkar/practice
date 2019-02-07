/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
*/

/*
4 ms
1.2 MB
*/
func reverse(x int) int {
    var res int
    for x != 0 {
        res = res * 10 + x%10
        x = x/10
    }
    if res < math.MinInt32 || res > math.MaxInt32 {
        return 0
    }
    return res
}