/*
We partition a row of numbers A into at most K adjacent (non-empty) groups, then our score is the sum of the average of each group.
What is the largest score we can achieve?

Note that our partition must use every number in A, and that scores are not necessarily integers.

Example:
Input:
A = [9,1,2,3,9]
K = 3
Output: 20
Explanation: 
The best choice is to partition A into [9], [1, 2, 3], [9]. The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
We could have also partitioned A into [9, 1], [2], [3, 9], for example.
That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.

Note:
1 <= A.length <= 100.
1 <= A[i] <= 10000.
1 <= K <= A.length.
Answers within 10^-6 of the correct answer will be accepted as correct.
*/

func largestSumOfAverages(A []int, K int) float64 {
    memo := make([][]float64, len(A))
    for i, _ := range memo {
        memo[i] = make([]float64, K)
        memo[i][0] = 0
    }

    return calculateAvg(A, 0, K, memo)
}

func calculateAvg(A []int, idx, K int, memo [][]float64) float64 {
    if memo[idx][K-1] > 0.0{
        return memo[idx][K-1]
    }

    if  K == 1  {
        memo[idx][K-1] = avgOf(A[idx:])
        return memo[idx][K-1]
    }

    max := 0.0
    for i:=idx;i<=len(A)-K;i++{
        if K-1 <= 0 || i+1 >= len(A){continue}
        sum := avgOf(A[idx:i+1]) + calculateAvg(A, i+1, K-1, memo)
        if sum > max {
            max = sum
        }
    }
    memo[idx][K-1] = max
    return memo[idx][K-1]
}

func avgOf(A []int) float64 {
    if len(A) == 0{
        return 0.0
    }

    var ans float64 
    for _, val := range A {
        ans += float64(val)
    }
    n := float64(len(A))
    return ans/n
}