package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	mid := (m + n) / 2
	var temp [2]int
	var result float64
	for k := 0; k <= mid; k++ {
		if nums1[i] <= nums2[j] {
			temp[0] = temp[1]
			temp[1] = nums1[i]
			i++
		} else {
			temp[0] = temp[1]
			temp[1] = nums2[j]
			j++
		}
	}
	if (m+n)%2 == 0 {
		result = (float64(temp[0]) + float64(temp[1])) / 2
	} else {
		result = float64(temp[1])
	}
	return result
}
