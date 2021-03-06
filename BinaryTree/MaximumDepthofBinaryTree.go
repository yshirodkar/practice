/*
Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if (root == nil) {
        return 0
    }

    ldepth := maxDepth(root.Left)
    rdepth := maxDepth(root.Right)
    if ldepth > rdepth {
        return ldepth + 1   
    } else {
        return rdepth + 1   
    }
}