/*
    x rank soldier report to x+1 rank officer. From given input array of ranks get the number soldier who can report to someone in same array
    input [1,2,5,7] and answer is 1
    1 can report to 2 but no other soldier can report to other.
*/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
    import "sort"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(ranks []int) int {
    // write your code in Go 1.4
    sort.Ints(ranks)
    rankMap := make(map[int]int)
    resp := 0
    i := len(ranks)-1
    for i >= 0 {
        if ok := rankMap[ranks[i] +1]; ok > 0 {
           resp++ 
        }
        rankMap[ranks[i]] = 1
        i--
    }
    return resp
}