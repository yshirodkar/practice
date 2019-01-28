/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var outPut []int
    var rStack []*TreeNode

    if root == nil {
        return outPut
    }
    for current := root; current != nil; {
        outPut = append(outPut, current.Val)
        if current.Left != nil {
            if current.Right != nil {
                rStack = append(rStack, current.Right)
            }
            current = current.Left
        } else {
            if current.Right != nil {
                current = current.Right
            } else {
                if len(rStack) == 0 {
                    break
                }
                current = rStack[len(rStack)-1]
                rStack = rStack[:len(rStack)-1]
            }
        }
    }
    return outPut
}