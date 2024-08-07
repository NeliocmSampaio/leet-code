package main

func main() {
	values1 := [4]int{}
	values2 := [4]int{}

	var nums1 []int
	var nums2 []int
	var m, n int

	nums1 = values1[0:]
	nums2 = values2[0:]
	m = 3
	n = 3

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1
	j := n - 1

	for index := m + n - 1; index >= 0; index-- {
		if i >= 0 && (j < 0 || nums1[i] > nums2[j]) {
			nums1[index] = nums1[i]
			i = i - 1
		} else {
			nums1[index] = nums2[j]
			j = j - 1
		}
	}
}
