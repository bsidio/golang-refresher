package leetcode

import "fmt"

func LoSuPal() string {
	str := "s"
	maxSize := 0
	sol := ""
	//BruteForce

	if len(str) == 1 {
		return str
	}

	for i := 0; i < len(str); i++ {
		if isPalindrome(str[:i]) {
			fmt.Println("isPla", str[:i])
			if len(str[:i]) > maxSize {
				fmt.Println("newsol ->", sol)
				sol = str[:i]
				maxSize = len(sol)

			}
		}
		fmt.Println("notPla ->", str[:i])
	}
	fmt.Println("final sol")
	return sol
}

func isPalindrome(s string) bool {
	fmt.Println("is", s)

	for i, j := 0, len(s)-1; i < len(s) && j > 0; i, j = i+1, j-1 {
		if s[i] != s[j] {
			fmt.Println("isf", s[i], s[j])
			return false
		}
	}

	return true
}
