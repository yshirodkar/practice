/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 
*/

/*
    Solution 1
*/
func twoSum(nums []int, target int) []int {
    ans := []int{}
    for ind, val := range nums {
        for ind_sec, val_sec := range nums {
            if (ind_sec > ind) {
                if (val_sec == target - val) {
                    ans = ans[:0]
                    ans = append(ans, ind)
                    ans = append(ans, ind_sec)
                }
            }
        }
    }
    return ans
}

/*
    Solution 2 (memory optimized and much faster)
*/
func twoSum(nums []int, target int) []int {
    ans := []int{}
    arrayLen := len(nums)
    if arrayLen > 0 {        
        for ind := 0; ind < arrayLen-1; ind++ {
            for ind_sec := ind+1;ind_sec < arrayLen; ind_sec++ {
                if (nums[ind_sec] == target - nums[ind]) {
                    ans = ans[:0]
                    ans = append(ans, ind)
                    ans = append(ans, ind_sec)
                }
            }
        }
    }
    return ans
}

/*
    Time optimized solution
*/
func twoSum(nums []int, target int) []int {
    seen := make(map[int] int)
    for idx, num := range(nums){
        if _ , ok := seen[num]; ok{
            return []int{seen[num], idx} 
        } else {
            seen[target - num] = idx
        }
    }
    return []int{0}
}