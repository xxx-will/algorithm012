package main

// n 加到m之后，sort
// 各自有序，可以反向遍历，从最后位置开始放  遍历 m，n中小的次数

func merge(nums1 []int, m int, nums2 []int, n int) {

	index := m + n - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}

	for j >= 0 {
		nums1[j] = nums2[j]
		j--
	}

}
