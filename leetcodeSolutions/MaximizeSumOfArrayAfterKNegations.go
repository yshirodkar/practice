/*
Given an array A of integers, we must modify the array in the following way:
we choose an i and replace A[i] with -A[i], and we repeat this process K times in total.
(We may choose the same index i multiple times.)

Return the largest possible sum of the array after modifying it in this way.

Example 1:
Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].

Example 2:
Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].

Example 3:
Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
*/
func largestSumAfterKNegations(A []int, K int) int {
    num := A
    sort.Ints(num)
    var i int
    for i = 0; i<K; i++ {
        if i > len(num)-1 || num[i] > 0 {
            break
        }
        num[i] = -1*num[i]
    }
    if K > i {
        if (K - i)%2 != 0 { // odd
            if i == 0 { // all postive value in A
                num[i] = -1 * num[i]
            } else if i >= len(num){ // all negative value in A
                num[i-1] = -1*num[i-1]
            } else {
                if num[i] > num[i-1] {
                    num[i-1] = -1 * num[i-1]
                } else {
                    num[i] = -1 * num[i]
                }
            }
        }
    }
    total := 0
    for i := 0; i < len(num);i++ {
        total += num[i]
    }
    return total
}