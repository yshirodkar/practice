func isUnique(s string) bool {
	tmpStr := make(map[rune]int)
	ss := []rune(s)
	i := 0
	for i < len(ss) {
		 _, ok := tmpStr[ss[i]]
		if ok {
			return false
		} else {
			tmpStr[ss[i]] = i
		}
		i++
	}
	return true
}