package median

func MedianSplit(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	lenz := len1 + len2
	var k int
	var extraStep bool
	if lenz%2 == 0 {
		k = lenz / 2
		extraStep = true
	} else {
		k = lenz/2 + 1
	}

	l1, l2, num := kth(nums1, nums2, k)
	if extraStep {
		_, _, num2 := kth(l1, l2, 1)
		return float64(num+num2) / 2.0
	} else {
		return float64(num)
	}
}

func kth(nums1 []int, nums2 []int, k int) (nums1_left, nums2_left []int, val int) {
	for {
		if len(nums1) == 0 {
			val = nums2[k-1]
			nums2 = nums2[k:]
			break
		}

		if len(nums2) == 0 {
			val = nums1[k-1]
			nums1 = nums1[k:]
			break
		}

		if k == 1 {
			if nums1[0] <= nums2[0] {
				val = nums1[0]
				nums1 = nums1[1:]
			} else {
				val = nums2[0]
				nums2 = nums2[1:]
			}

			break
		}

		split := k / 2

		var fetch1, fetch2 int
		if split >= len(nums1) {
			fetch1 = len(nums1)
			fetch2 = k - fetch1
		} else if split >= len(nums2) {
			fetch2 = len(nums2)
			fetch1 = k - fetch2
		} else {
			fetch1 = split
			fetch2 = split
		}

		if nums1[fetch1-1] < nums2[fetch2-1] {
			nums1 = nums1[fetch1:]
			k -= fetch1
			val = nums2[fetch2-1]
		} else {
			nums2 = nums2[fetch2:]
			k -= fetch2
			val = nums1[fetch1-1]
		}
	}

	nums1_left = nums1
	nums2_left = nums2

	return
}
