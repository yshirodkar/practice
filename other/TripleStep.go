/*
A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3
steps at a time. Implement a method to count how many possible ways the child can run up the
stairs. 
*/

package main

import "fmt"

func main() {
	stepCount := 4
	comb := tripleStep(stepCount)
	fmt.Println("Step combination is ", comb)
}

func tripleStep (stepCount int) int{

	if stepCount < 0 {
		return 0
	} else if stepCount == 0 {
		return 1
	} else {
		return tripleStep(stepCount - 1) + tripleStep(stepCount - 2) + tripleStep(stepCount - 3)
	}
}