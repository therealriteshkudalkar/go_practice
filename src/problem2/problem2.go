package problem2

func isPalindromic(str string) bool {
	n := len(str)
	i := 0
	j := n - i - 1
	for i < j {
		if str[i:i+1] != str[j:j+1] {
			return false
		}
		i++
		j = n - i - 1
	}
	return true
}

func LongestPalindrome(s string) string {
	// Find all the substrings and check if it is a palindrome and
	palindromicStrings := map[string]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			currentString := s[i:j]
			if isPalindromic(currentString) {
				palindromicStrings[currentString] = j - i
			}
		}
	}
	lengthOfLargestPalindromicString := 0
	largestPalindromicString := ""
	for palindromicString, lengthOfCurrentPalindromicString := range palindromicStrings {
		if lengthOfCurrentPalindromicString > lengthOfLargestPalindromicString {
			lengthOfLargestPalindromicString = lengthOfCurrentPalindromicString
			largestPalindromicString = palindromicString
		}
	}
	return largestPalindromicString
}
