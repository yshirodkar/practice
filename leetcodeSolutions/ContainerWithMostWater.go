/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

/*
568 ms
4.6MB
*/
func maxArea(height []int) int {
    maxArea := 0
    maxheight := 0
    for i := 0; i < len(height) -1; i++ {
        for j := i+1; j < len(height); j++ {
            if height[i] >= height[j] {
                maxheight = height[j]
            } else {
                maxheight = height[i]
            }
            if maxArea < (maxheight * (j-i)) {
                maxArea = (maxheight * (j-i))
            }
        }
    }
    return maxArea
}