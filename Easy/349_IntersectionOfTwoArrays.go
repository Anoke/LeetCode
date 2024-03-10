package Easy

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i := 0
	j := 0
	var ans []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if len(ans) == 0 || ans[len(ans)-1] != nums1[i] {
				ans = append(ans, nums1[i])
			}
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}
