// https://leetcode.cn/problems/median-of-two-sorted-arrays/
/*
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10**6 <= nums1[i], nums2[i] <= 10**6
*/

// 不是最优解
package median

func Median(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	var dst int
	var extraStep bool
	lenz := len1 + len2
	if lenz%2 == 0 {
		dst = lenz / 2
		extraStep = true
	} else {
		dst = lenz/2 + 1
	}

	step1 := -1
	step2 := -1
	steps := 0
	want := 0
	extra := 0

	can1 := true
	can2 := true
	for {
		if steps == dst+1 {
			return float64(extra+want) / 2.0
		}

		if steps == dst {
			if !extraStep {
				return float64(want)
			} else {
				extra = want
			}
		}

		if can1 && step1+1 >= len1 {
			can1 = false
		}

		if can2 && step2+1 >= len2 {
			can2 = false
		}

		if !can1 {
			step2 += 1
			want = nums2[step2]
			steps += 1
			continue
		}

		if !can2 {
			step1 += 1
			want = nums1[step1]
			steps += 1
			continue
		}

		if nums1[step1+1] < nums2[step2+1] {
			step1 = step1 + 1
			want = nums1[step1]
		} else {
			step2 = step2 + 1
			want = nums2[step2]
		}
		steps += 1
	}
}
