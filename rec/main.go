package main

import "fmt"

func fib(n int, dp []int) int {
	if n <= 1 {
		return 0
	}

	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = fib(n-1, dp) + fib(n-2, dp)
	return dp[n]
}

func main() {
	var n int
	fmt.Printf("Enter the value of %d: ", n)
	fmt.Scanf("%d", &n)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	fmt.Printf("The fibonacci of %d is : %d", n, fib(n, dp))

}
