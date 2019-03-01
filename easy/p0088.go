package easy

func Merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, 0, m+n)
	max := m + n
	i1, i2 := 0, 0
	for i := 0; i < max; i++ {
		if i1 >= m || i2 >= n {
			break
		}
		if nums1[i1] <= nums2[i2] {
			tmp = append(tmp, nums1[i1])
			i1++
			continue
		} else {
			tmp = append(tmp, nums2[i2])
			i2++
		}
	}
	if i1 < m {
		tmp = append(tmp, nums1[i1:m]...)
	} else if i2 < n {
		tmp = append(tmp, nums2[i2:n]...)
	}
	copy(nums1, tmp)
}
