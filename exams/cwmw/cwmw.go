// 11. Container With Most Water

package cwmw

func ContainWithMostWater(nums []int) (res int) {
	i := 0
	j := len(nums) - 1

	for j > i {
		res = maxInt(res, minInt(nums[i], nums[j])*(j-i))

		if nums[i] < nums[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func minInt(x, y int) int {
	if x > y {
		return y
	}

	return x
}
