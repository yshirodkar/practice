/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the
other. 
*/
package main
import  "fmt"
func main () {
	tmps1 := "god"
	tmps2 := "dog"
	res := isPermutation(tmps1, tmps2)
	fmt.Println("Res is ", res)
}

/*
	Dog = god
*/
func isPermutation (s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	stringMap := make(map[rune]int)
	str1 := []rune(s1)
	i := 0
	for i < len(s1) {
		_, ok :=  stringMap[str1[i]]
		if ok {
			stringMap[str1[i]] = stringMap[str1[i]] + 1
		} else {
			stringMap[str1[i]] = 1
		}
		i++
	}

	str2 := []rune(s2)
	j := 0
	for j < len(str2) {
		_, ok := stringMap[str2[j]]
		if ok && stringMap[str2[j]] >0 {
			stringMap[str2[j]] = stringMap[str2[j]] - 1
		} else {
			return false
		}
		j++
	}
	return true
}
