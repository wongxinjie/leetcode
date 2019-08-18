/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/17
*/
package p279

func numSquares(n int) int {
	dp:= make([]int, 0)
	dp = append(dp, 0)

	for i := 1; i <= n; i++ {
		dp = append(dp, i)
		for j := 1; j * j <= i; j++ {
			if j * j == i {
				if dp[i] > 1 {
					dp[i] = 1
				}
				break
			} else {
				if dp[i] > dp[i - j * j] + 1 {
					dp[i] = dp[i - j * j] + 1
				}
			}
		}
	}
	return dp[n]
}
