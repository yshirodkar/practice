type Node struct {
	val   int
	left  *Node
	right *Node
}
 
func PreorderRecursive(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.val)
		PreorderRecursive(root.left)
		PreorderRecursive(root.right)
	}
}