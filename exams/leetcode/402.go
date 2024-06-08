package leetcode

/*
给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。


示例 1 ：

输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
示例 2 ：

输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 ：

输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。


提示：

1 <= k <= num.length <= 105
num 仅由若干位数字（0 - 9）组成
除了 0 本身之外，num 不含任何前导零
*/

func removeKdigits(num string, k int) string {
	k = len(num) - k
	if k == 0 {
		return "0"
	}

	stack := []rune{rune(num[0])}

	for i, digit := range num {
		if i == 0 {
			continue
		}

		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j] > digit {
				// 后面的数字只能填满 stack 了，不再进行比较
				if k-len(stack) >= len(num)-i-1 {
					break
				}

				stack[j] = digit
				stack = stack[0 : j+1]
			} else {
				if len(stack) < k {
					stack = append(stack, digit)
				}
				break
			}
		}
	}

	// stack 未满， 从 num 末端补齐
	if len(stack) < k {
		stack = append(stack, []rune(num[len(num)-len(stack)-1:])...)
	}

	// remove foremost 0
	for i, r := range stack {
		if string(r) != "0" {
			stack = stack[i:]
			break
		} else {
			// 全是 0 的情况
			if i == len(stack)-1 {
				return "0"
			}
		}
	}

	return string(stack)
}
