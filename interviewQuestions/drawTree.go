/*
Given height of tree and boolean value which decides to fill the tree or not. Draw a tree
e.g 
drawtree(6, true)
    X    
   XXX   
  XXXXX  
 XXXXXXX 
XXXXXXXXX
    X  

drawtree(6, false)
    X    
   X X   
  X   X  
 X     X 
XXXXXXXXX
    X  
*/
package main

import "fmt"

func main() {
	drawtree(6, true)
}

func drawtree(height int, fill bool) {

	width := getWidth(height-1)
	var start = width/2
	var end = width/2
	for j := 0; j< height-2; j++ {
		var str = ""
		for i := 0; i<width; i++ {
			if i != start && i != end {
				if i < start && i > end {
					str = str + " "
				} else if i > start && i < end && fill {
					str = str + "X"
				} else {
					str = str + " "
				}
			}
			if i == start && i != end {
				str = str + "X"
			}
			if i == end && i != start {
				str = str + "X"
			}
			if i == start && i == end {
				str = str + "X"
			}
		}
		start = start - 1
		end = end +1
		fmt.Println(str)
	}
	lastRow(width)
	stem(width)
}

func getWidth(num int) int {
	var width = 1
	for i := 0; i < num-1; i++ {
		width = width + 2
	}
	return width
}

func stem(num int) {
	var str = ""
	for i:=0; i< num; i++ {
		if i == num/2  {
			str = str + "X"
		} else {
			str = str + " "
		}
	}
	fmt.Println(str)
}
func lastRow(num int) {
	var str = ""
	for i:=0; i< num; i++ {
		str = str + "X"
	}
	fmt.Println(str)
}