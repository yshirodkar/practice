/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/
func generate(numRows int) [][]int {
    s := make([][]int, numRows)
    for i:= 0; i < numRows; i++ {
        s[i] = make([]int, i + 1)
        s[i][0] = 1
        s[i][i] = 1
        for j := 1; j < i; j++ {
            s[i][j] = s[i-1][j] + s[i-1][j-1]
        }
    }
    return s
}