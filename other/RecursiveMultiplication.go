/*
 Write a recursive function to multiply two positive integers without using the
* operator. You can use addition, subtraction, and bit shifting, but you should minimize the number
of those operations.

*/
func minProduct(a, b int) int {
	bigger := a 
	if a < b {
		bigger = b
	}
	smaller := b
	if a < b {
		smaller = a
	}
	return minProductHelper(smaller, bigger)
}

func minProductHelper(smaller, bigger int) int {
	if smaller == 0{
 		return 0
	} else if smaller == 1 {
		return bigger
	}

	s := smaller >> 1
	halfProd := minProductHelper(s, bigger)
	if smaller%2 == 0 {
		return halfProd + halfProd
	} else {
		return halfProd + halfProd + bigger
	}
}