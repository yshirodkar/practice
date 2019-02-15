/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

/*
First define your boundries. As you reach to boundry reduce that respective border.
There are four cases
    1. Move right (col++)
    2. Move down (row++)
    3. Move left (col--)
    4. Move top (row--)
All these need to be in loop
*/
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<Integer>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return res;
        }
        int left = 0;
        int right = matrix[0].length - 1;
        int top = 0;
        int bottom = matrix.length - 1;
        int tag = 0;
        int row = top;
        int col = left;
        // int m = 0;
        for (int m = 0; m < matrix.length * matrix[0].length; m++) {
            if (tag == 0) {
                res.add(matrix[row][col]);
                if (col == right) {
                    tag = 1;
                    top++;
                    row = top;
                } else {
                    col++;
                }
            } else if (tag == 1) {
                res.add(matrix[row][col]);
                if (row == bottom) {
                    tag = 2;
                    right--;
                    col = right;
                } else {
                    row++;
                }
            } else if (tag == 2) {
                res.add(matrix[row][col]);
                if (col == left) {
                    tag = 3;
                    bottom--;
                    row = bottom;
                } else {
                    col--;
                }
            } else {
                res.add(matrix[row][col]);
                if (row == top) {
                    tag = 0;
                    left++;
                    col = left;
                } else {
                    row--;
                }
            }
        }
        return res;
    }
}