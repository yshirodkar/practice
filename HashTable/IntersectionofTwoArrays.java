/*
Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:
Each element in the result must be unique.
The result can be in any order.
*/

class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        Set<Integer> set = new HashSet<>();
        Set<Integer> intersection = new HashSet<>();
        for(int i = 0; i< nums1.length; i++) {
            set.add(nums1[i]);
        }
        
        for (int i = 0; i< nums2.length; i++) {
            if (set.contains(nums2[i])) {
                if (!intersection.contains(nums2[i])) {
                    intersection.add(nums2[i]);   
                }
            }
        }
        
        int[] res = new int[intersection.size()];
        int i = 0;
        for(int item : intersection) {
            res[i] = item;
            i++;
        }
        return res;
    }
}