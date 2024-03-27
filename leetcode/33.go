package leetcode

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为
[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例
如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则
返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0输出：4示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3输出：-1示例 3：

输入：nums = [1], target = 0输出：-1


提示：

1 <= nums.length <= 5000 -10**4 <= nums[i] <= 10**4 nums 中的每个值都 独一无二题目数据保证 nums 在预先未
知的某个下标上进行了旋转 -10**4 <= target <= 10**4
*/

func search(nums []int, target int) int {
	var f func([]int, int, int) int
	f = func(sortedNums []int, t, baseIndex int) int {
		if len(sortedNums) == 1 {
			if t == sortedNums[0] {
				return baseIndex
			} else {
				return -1
			}
		}

		middleIdx := len(sortedNums) / 2
		if sortedNums[middleIdx] > t {
			return f(sortedNums[:middleIdx], t, baseIndex)
		} else if sortedNums[middleIdx] == t {
			return baseIndex + len(sortedNums)/2
		} else {
			if len(sortedNums) == 2 {
				// 后半段是空的
				return -1
			} else {
				return f(sortedNums[middleIdx+1:], t, baseIndex+middleIdx+1)
			}
		}
	}

	var g func([]int, int) int
	g = func(numbers []int, baseIndex int) int {
		if len(numbers) == 1 {
			if numbers[0] == target {
				return baseIndex
			} else {
				return -1
			}
		}

		// 本身有序, 即翻转次数等于数组长度
		if numbers[0] < numbers[len(numbers)-1] {
			ii := f(numbers, target, 0)
			if ii == -1 {
				return ii
			} else {
				return baseIndex + ii
			}
		}

		// 翻转了
		middleIdx := len(numbers)/2 - 1

		if middleIdx == 0 {
			// 前半段只有一个
			if numbers[middleIdx] == target {
				return baseIndex
			} else {
				return g(numbers[1:], baseIndex+1)
			}
		}

		if numbers[0] < numbers[middleIdx] {
			// 前半段有序

			// 如果在前半段范围内
			if numbers[0] <= target && target <= numbers[middleIdx] {
				ii := f(numbers[0:middleIdx+1], target, 0)
				if ii == -1 {
					return ii
				} else {
					return baseIndex + ii
				}
			} else {
				// 在后半段寻找

				ii := g(numbers[middleIdx+1:], middleIdx+1)
				if ii == -1 {
					return ii
				} else {
					return baseIndex + ii
				}
			}
		} else {
			// 前半段无序

			// 如果在后半段范围内
			if numbers[middleIdx+1] <= target && target <= numbers[len(numbers)-1] {
				ii := f(numbers[middleIdx+1:], target, middleIdx+1)
				if ii == -1 {
					return -1
				} else {
					return baseIndex + ii
				}
			}

			// 在前半段寻找
			ii := g(numbers[:middleIdx+1], 0)
			if ii == -1 {
				return ii
			} else {
				return baseIndex + ii
			}
		}
	}

	return g(nums, 0)
}
