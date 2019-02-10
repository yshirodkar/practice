/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
Input: [2,2,1]
Output: 1

Example 2:
Input: [4,1,2,1,2]
Output: 4

*/

func singleNumber(nums []int) int {
    tmp := make(map[int]int)
    for _, num := range nums {
        tmp[num] = tmp[num] + 1
    }
    var nonDuplicate int
    for ind, val := range tmp {
        if val < 2 {
            nonDuplicate = ind
            break
        }
    }
    return nonDuplicate
}