package main

func minPathSum(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	columns := len(grid[0])

	// 状态空间，走到dp[i][j] 最短的路径长
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	// dp状态的初始化
	dp[0][0] = grid[0][0]

	//第一行，只有左边
	for i := 1; i < columns; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for j := 1; j < rows; j++ {
		dp[j][0] = dp[j-1][0] + grid[j][0]
	}
	//

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][columns-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
