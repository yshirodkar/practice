/*
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deepCopyWithAdd(root *TreeNode, v int) *TreeNode {
    if root == nil { return nil }
    res := &TreeNode{root.Val+v, nil, nil}
    res.Left, res.Right = deepCopyWithAdd(root.Left, v), deepCopyWithAdd(root.Right, v)
    return res
}

func generateTrees(n int) []*TreeNode {
    if n<=0 { return []*TreeNode{}}
    l := [][]*TreeNode{}
    l = append(l, []*TreeNode{nil})
    l = append(l, []*TreeNode{&TreeNode{1, nil, nil}})
    for i:=2; i<=n; i++ {
        l2 := []*TreeNode{}
        for j:=0; j<i; j++ {
            for _, left := range l[j] {
                for _, right := range l[i-j-1] {
                    parent := &TreeNode{j+1, deepCopyWithAdd(left, 0), deepCopyWithAdd(right, j+1)}
                    l2 = append(l2, parent)
                }
            }
        }
        l = append(l, l2)
    }
    return l[n]
}