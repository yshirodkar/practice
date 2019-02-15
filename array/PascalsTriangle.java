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

class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();
        for (int i = 0; i < numRows; i ++) {
            List<Integer> list = new ArrayList<Integer>(Arrays.asList(1));
            for (int j=1; j<i; j++) {
                list.add(res.get(i-1).get(j-1)+res.get(i-1).get(j));
            }
            if (i>0) {
                list.add(1);
            }
            res.add(list);
        }
        return res;
    }
}