/*
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/
func topKFrequent(nums []int, k int) []int {
    tmpMap := make(map[int]int)

    for _, num := range nums {
        if tmpMap[num] > 0 {
            tmpMap[num] += 1
        } else {
            tmpMap[num] = 1
        }
    }
    
    myNums := make(MyNums, len(tmpMap))
	i := 0
	for num, count := range tmpMap {
		myNums[i] = MyNum{Val: num, Count: count}
		i++
	}

	heap.Init(&myNums)
	var res []int
	for i := 0; i < k; i++ {
		num := heap.Pop(&myNums).(MyNum)
		res = append(res, num.Val)
	}
    return res
}

// MyNum stores its value and frequency as Count.
type MyNum struct {
	Val   int
	Count int
}

type MyNums []MyNum

func (n MyNums) Len() int {
	return len(n)
}

func (n MyNums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n MyNums) Less(i, j int) bool {
	return n[i].Count >= n[j].Count
}

func (n *MyNums) Push(num interface{}) {
	myNum := num.(MyNum)
	*n = append(*n, myNum)
}

func (n *MyNums) Pop() interface{} {
	tmp := *n
	l := len(tmp)
	var res interface{} = tmp[l-1]
	*n = tmp[:l-1]
	return res
}
