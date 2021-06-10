//704. Binary Search

package binarysearch

func BinarySearch(nums []int, target int) int {
	l := len(nums)
	leftMargin := 0
	rightMargin := l - 1

	for i := l / 2; i != 0 || i != l; {
		if target == nums[i] {
			return i
		}

		if target < nums[i] {
			rightMargin = i - 1
		} else {
			leftMargin = i + 1
		}

		i = (leftMargin + rightMargin) / 2

		if leftMargin > rightMargin {
			return -1
		}
	}

	return -1
}
