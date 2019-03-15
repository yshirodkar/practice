/*
For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.

 

Example 1:

Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
Example 2:

Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
Example 3:

Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
Example 4:

Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000
 

Note：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
If A.length > 1, then A[0] != 0
*/
func addToArrayForm(A []int, K int) []int {
    carry := 0
    aLen := len(A)
    tmRes := make([]int, aLen)
    for i:=aLen-1; i >=0; i-- {
        tmpVal := A[i] + K%10 + carry
        if tmpVal > 9 {
            carry = 1
        } else {
            carry = 0
        }
        tmRes[i] = tmpVal%10
        K = K/10
    }

    for K > 0 {
        tmpVal := K%10 + carry
        if tmpVal > 9 {
            carry = 1
        } else {
            carry = 0
        }
        tmRes = append([]int{tmpVal%10}, tmRes...)
        K = K/10
    }
    if carry == 1 {
        return append([]int{1}, tmRes...)
    }
    return tmRes
}