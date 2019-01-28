/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var stack []*TreeNode
    var res []int
    current := root
    for current != nil {
        if current.Left == nil && current.Right == nil {
            res = append(res, current.Val)
            if len(stack) == 0 {
                break
            }
            current = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        } else {
            left, right := current.Left, current.Right
            current.Left, current.Right = nil, nil
            stack = append(stack, current)
            if left != nil {
                current = left
                if right != nil {
                    stack = append(stack, right)
                }
            } else {
                current = right
            }
        }
    }
    return res
}