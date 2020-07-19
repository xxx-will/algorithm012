/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) //key:  num  value: index
	for i, n := range nums {
		other := target - n
		if j, ok := m[other]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

	return []int{}
}

// @lc code=end

