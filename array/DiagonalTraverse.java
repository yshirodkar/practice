/*
Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

Example:
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

Output:  [1,2,4,7,5,3,6,8,9]

Explanation:

Note:
The total number of elements of the given matrix will not exceed 10,000.
*/

public class Solution {
    public int[] findDiagonalOrder(int[][] matrix) {
        if (matrix == null || matrix.length == 0)   return new int[0];
        int m = matrix.length, n = matrix[0].length;
        int[] res = new int[m*n];
        
        for (int layer = 0, i = 0; layer < m+n-1; layer++) {
            int upperBoundx = Math.min(layer, m-1);
            int lowerBoundx = Math.max(0, layer - n + 1);
            
            if (layer % 2 != 0) {
                for (int x = lowerBoundx; x <= upperBoundx; x++) {
                    res[i++] = matrix[x][layer-x];
                }
            } else {
                
                for (int x = upperBoundx; x >= lowerBoundx; x--) {
                    res[i++] = matrix[x][layer-x];
                }
            }
        }
        return res;
    }
}