package main

// 循环数组

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}

	k = k % l

	cnt := 0
	pre := nums[0]
	start := 0
	cur := start
	for cnt < l {

		next := (k + cur) % l

		tmp := nums[next]
		nums[next] = pre
		pre = tmp

		cnt++
		cur = next
		if start == next && cnt < l {
			start++
			cur = start
			pre = nums[start]
		}
	}

}
