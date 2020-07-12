package main

import "fmt"

// 删除有序数组重复元素
func delDup(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	j := 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[j] {
			j++
			if i > j {
				nums[j] = nums[i]
			}
		}
	}

	return j + 1
}

func main() {
	test := []int{1, 2, 3, 4, 4, 4, 5}
	l := delDup(test)

	fmt.Println("len = ", l)

	test = []int{1, 1, 1, 1, 1, 1, 1}
	l = delDup(test)

	fmt.Println("len = ", l)

	test = []int{1, 2, 2, 2, 4, 4, 4, 5}
	l = delDup(test)

	fmt.Println("len = ", l)

	hello := make([]int, 0)
	hello = append(hello, 1)

	fmt.Println("cap= ", cap(hello))

	wold := [10]int{}
	fmt.Println("wold ", wold[9])

}
