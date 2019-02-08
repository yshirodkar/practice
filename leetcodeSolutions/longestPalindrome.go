func longestPalindrome(s string) string {
	lenn:=len(s)//palindromic substring
	if lenn==0||lenn==1{
		return s
	}
	var longest [2]int
	for i:=0;i<2*lenn-1;i++{
		var start_end [2]int
		fmt.Println("Starting expand ", i)
		start_end=expand(i,s)
		if start_end[1]-start_end[0]>longest[1]-longest[0]{
			longest[1]=start_end[1]
			longest[0]=start_end[0]
		}
	}
	return s[longest[0]:longest[1]+1]
}
func expand(i int,s string)[2]int{
	var start_end [2]int
	var left int
	var right int
	lenn:=len(s)
	if i==0||i==2*lenn-2{
		fmt.Println(i)
		start_end[0]=i
		start_end[1]=i
		return start_end
	}
	if i%2==0{
		left=i/2 - 1
		fmt.Println("left ", left)
		right=i/2 + 1
		fmt.Println("right1 ", right)
		right=1+(i+1)/2
		fmt.Println("right2 ", right)
	}else{
		left=(i-1)/2
		fmt.Println("left2 ", left)
		right=(i+1)/2
		fmt.Println("right3 ", right)
	}
    for left>-1&&right<lenn{
    	fmt.Println("left3 ", left)
    	fmt.Println("right4 ", right)
			if left==0||right==lenn-1{
				if s[left]==s[right]{
					start_end[0]=left
					start_end[1]=right
					break
				}
			}
			if s[left]!=s[right]{
				start_end[0]=left+1
				start_end[1]=right-1
				fmt.Println("break 2 ", start_end)
				break
			}
			fmt.Println("start_end ", start_end)
			left--
			right++
		}
	return start_end
}