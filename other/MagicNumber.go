/*
	nums = [1,2,2,3,3,8,9,11]
*/
func hasMagicNumber(nums []int, start, end int) bool {
	if end < start {
		return false
	}

	mid := (start + end) /2
	if nums[mid] == mid {
		return true
	} else if nums[mid] > mid {
		return hasMagicNumber(nums, start, mid -1)
	} else {
		return hasMagicNumber(nums, mid +1, end)
	}
}