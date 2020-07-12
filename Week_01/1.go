package main

//
// 先排序 双指针 两头移动 ： 不行。输出为 index
// hash map

func twoSum(nums []int, target int) []int {

	indexs := []int{0, 0}
	m := make(map[int]int)

	for i, n := range nums {
		pair := target - n
		if v, ok := m[pair]; ok {
			indexs[0] = v
			indexs[1] = i
			break
		}

		m[n] = i
	}

	return indexs
}
