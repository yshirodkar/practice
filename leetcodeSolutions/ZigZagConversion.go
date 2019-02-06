/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/

/*
8 ms
3.1 MB
*/
func convert(s string, numRows int) string {
  if numRows == 1 {
    return s
  }
  bytes := make([]byte, len(s))
  index := 0
  cycleLen := numRows * 2 - 2
  for i := 0; i < numRows; i++ {
    for j := 0; j + i < len(s); j = j + cycleLen {
      bytes[index] = s[i + j]
      index++
      if (i != 0 && i != numRows - 1 && j + cycleLen - i < len(s)) {
        bytes[index] = s[j + cycleLen - i]
        index++
      }
    }
  }
  return string(bytes)
}