/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}

	return helper(0, 1, 2, N)

}

func helper(a, b, level, N int) int {

	if level == N {
		return a + b
	}
	// logic

	tmp := b
	b = a + b
	a = tmp

	// drill-down
	return helper(a, b, level+1, N)

}

// @lc code=end

