/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:
Input: [3,2,1,5,6,4] and k = 2
Output: 5

Example 2:
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/
func findKthLargest(nums []int, k int) int {
    k = len(nums) - k + 1
    for {
        idx := quickSelect(nums)
        if k == idx + 1 {
            return nums[idx]
        } else if k > idx + 1 {
            nums = nums[idx + 1:]
            k = k - idx - 1
        } else {
            nums = nums[:idx]
        }
    }
    return -1
}

func quickSelect(n []int) int {
    n[0], n[len(n) - 1] = n[len(n) - 1], n[0]
    next := 0
    for i := 0; i < len(n) - 1; i++ {
        if n[i] < n[len(n) - 1] {
            n[i], n[next] = n[next], n[i]
            next++
        }
    }
    n[next], n[len(n) - 1] = n[len(n) - 1], n[next]
    return next
}