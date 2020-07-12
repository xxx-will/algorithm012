package main

//双指针起点都是0
//前面全是0的时候, j==i 自身交换。 不需要操作
//为了循环， i++,  被交换了后j++

// j 表示最后非0 元素的位置
// j 之前所有的元素非0
// j 和i 之间全是0

func moveZeroes(nums []int) {

	if len(nums) == 0 {
		return
	}

	j := 0
	for i, n := range nums {
		if n != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
