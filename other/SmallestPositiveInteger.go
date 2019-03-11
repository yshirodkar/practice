/*
Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
Copyright 2009–2019 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
*/

func Solution(A []int) int {
    // write your code in Go 1.4
    arrayLength := len(A)
    if arrayLength > 0 {
        sort.Ints(A)
        if A[arrayLength-1] <=0 {
            return 1
        } else {
            return getSmallestPositive(A)
        }
    } else {
        return 1
    }
    
}

func getSmallestPositive(arr []int) int {
      smallestPositive := 1
      for i :=0 ; i< len(arr); i ++ {
          if arr[i] == smallestPositive {
              smallestPositive = smallestPositive +1
          }
      }
      return smallestPositive
}