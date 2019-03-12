package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    cutDownCount := 0
    for ind, _ := range A {
        
        treesArr := cutTree(A, ind)
        if isInOrderCheck(treesArr) {
            cutDownCount++
        }
    }
    return cutDownCount
}

/*
    Remove the tree from Array
*/
func cutTree (trees []int, rmTree int) []int {
    var res []int
    for ind, tree := range trees {
        if ind != rmTree {
            res = append(res, tree)
        }
    }
    return res
}

/*
    Check trees are in non-descreasing order
*/
func isInOrderCheck(trees []int) bool {
    var valShouldBe int
    for ind, tree := range trees {
        if ind == 0 {
            valShouldBe = tree
        }
        if tree >= valShouldBe {
            valShouldBe = tree
        } else {
            return false
        }
    }
    return true
}