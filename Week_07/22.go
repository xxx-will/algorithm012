package main

import "fmt"

// 状态空间
// dp[i], 第i组括号生成的
// 状态方程：
// dp[i] = "(" + dp[j] + ")" + dp[i-1-j]  // 插入dp[i-1]中 插入位置 包含 第j对括号
func generateParenthesis(n int) []string {

	dp := make([][]string, 0)

	s := []string{""}
	dp = append(dp, s)

	for i := 1; i <= n; i++ {
		ss := make([]string, 0)
		for j := 0; j < i; j++ {
			for _, str1 := range dp[j] {
				for _, str2 := range dp[i-1-j] {
					str := "(" + str1 + ")" + str2
					ss = append(ss, str)
				}
			}
		}
		dp = append(dp, ss)
		fmt.Println(dp)
	}

	return dp[n]
}
