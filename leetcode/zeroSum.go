package leetcode

//Given an integer n, return any array containing n unique integers such that they add up to 0.

func ZeroSum(n int) []int {
	solution := make([]int, n)

	for i := 0; i < n/2; i++ {
		solution[i+1] = i + 1
		solution[n-i-1] = -(i + 1)
	}
	return solution

}
