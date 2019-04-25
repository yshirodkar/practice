/*
Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Example 1:

Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2
Example 2:

Input: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
Follow up:

A solution using O(n) space is pretty straight forward.
Could you devise a constant space solution?
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Record struct {
    first, second, prev *TreeNode
}

func inorder(t *TreeNode, r *Record) {
    if t == nil {
        return
    }
     
    inorder(t.Left, r)
    if r.prev != nil && r.prev.Val > t.Val {
        if r.first == nil {    
            r.first, r.second = r.prev, t
        } else {    
            r.second = t
        }
    }
    r.prev = t

    inorder(t.Right, r)
}

func recoverTree(root *TreeNode)  {
        r := &Record{}
    inorder(root, r)
    
    if r.first != nil {
        r.first.Val, r.second.Val = r.second.Val, r.first.Val
    }
}