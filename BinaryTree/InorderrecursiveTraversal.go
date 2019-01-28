/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }

    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }

    outPut := inorderTraversal(root.Left)
    outPut = append(outPut, root.Val)
    outPut = append(outPut, inorderTraversal(root.Right)...)
    return outPut
}