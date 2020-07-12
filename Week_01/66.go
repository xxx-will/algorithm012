package main

// 重复点，看是否进位，如果计算当前位刚好加1
// 扩展，如果不是加1，加其他数，则 最后一位(循环开始特殊处理) len()-2 开始重复

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i++ {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}

	new := make([]int, len(digits)+1)
	new[0] = 1

	return new
}
