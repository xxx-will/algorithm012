package main

func longestValidParentheses(s string) int {

	dp := make([]int, len(s))
	longest := 0

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else { // ....)) 形式
				// 找前一半的 '('
				pair := i - dp[i-1] // i-dp[i-1] 防止数组越界
				if pair > 0 && s[pair-1] == '(' {
					if pair >= 2 {
						dp[i] = dp[i-1] + dp[pair-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}

			longest = max(longest, dp[i])
		}
	}

	return longest
}
