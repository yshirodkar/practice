/*
Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.

Example 1:
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output: 
"apple"
Example 2:
Input:
s = "abpcplea", d = ["a","b","c"]

Output: 
"a"
Note:
All the strings in the input will only contain lower-case letters.
The size of the dictionary won't exceed 1,000.
The length of all the strings in the input won't exceed 1,000.
*/

func findLongestWord(s string, d []string) string {
    maxLen := 0
    var maxword string
    for _, word := range d {
        isword := true
        tmpS := s
        tmp := strings.Split(word, "")
        tmplen := len(tmp)
        i := 0
        for i < tmplen {
            if strings.Index(tmpS, tmp[i]) >=0 {
                tmpS = tmpS[:strings.Index(tmpS, tmp[i])] + tmpS[strings.Index(tmpS, tmp[i])+1 :]
            } else {
                isword = false
                break
            }
            i++
        }

        if (isword) {
            if (maxLen <= i) {
                if i == maxLen {
                    if maxword != word {
                        if maxword > word {
                            maxword = word
                        }
                    }
                } else if maxword < word {
                    maxword = word
                }
                maxLen = i
            }
        }
    }
    return maxword
}