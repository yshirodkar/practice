/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
func addStrings(num1 string, num2 string) string {
	str1, str2 := []byte(num1), []byte(num2)
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	sum := byte(0)
	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0; i, sum = i-1, sum/10 {
		if j >= 0 {
			sum += str2[j] - '0'
			j--
		}
		sum += str1[i] - '0'
		str1[i] = sum%10 + '0'
	}

	if (sum) != byte(0) {
		str1 = append([]byte{'1'}, str1...)
	}
	return string(str1)
}