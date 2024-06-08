package od

/*
题目描述

给定一个字符串 s，最多只能进行一次变换，返回变换后能得到的最小字符串(按照字典 序进行比较)变换规则:
交换字符串中任意两个不同位置的字符

输入描述

一串小写字母组成的字符串 s

输出描述

按照要求进行变换得到的最小字符串。

用例1

输入abcdef输出abcdef说明abcdef已经是最小字符串，不需要交换。用例2

输入bcdefa输出acdefb说明a和b进行位置交换，可以得到最小字符串
*/

// 递归，极限情况下，效率极低，如 aaaabbbccccddddefghijklmnzy, 只有交换最后2位
func minString(input string) string {
	if len(input) < 2 {
		return input
	}

	if string(input[0]) == "a" {
		return "a" + minString(input[1:])
	}

	minIndex := 0
	minChar := string(input[minIndex])
	changed := false

	for i := 1; i < len(input); i++ {
		c := string(input[i])
		if c <= minChar {
			if c < minChar {
				changed = true
			}
			minIndex = i
			minChar = c
		}
	}

	if changed {
		return string(input[minIndex]) + input[1:minIndex] + string(input[0]) + input[minIndex+1:]
	} else {
		return input[:minIndex+1] + minString(input[minIndex+1:])
	}
}

func minString2(input string) string {
	occurences := make([][]int, 26)

	var base rune = 'a'
	// record all occurence of every char
	for i, c := range input {
		idx := c - base
		if len(occurences[idx]) > 0 {
			occurences[idx] = append(occurences[idx], i)
		} else {
			occurences[idx] = []int{i}
		}
	}

	smallestLastIndex := 0
	baseIndex := 0
	toSwitchIndex := 0
	for _, occurence := range occurences {
		if len(occurence) == 0 {
			continue
		}

		smallestLastIndex = occurence[len(occurence)-1]
		if smallestLastIndex == baseIndex+len(occurence)-1 {
			// 如果前面的全是这个最小字符，查看下个字符
			baseIndex = smallestLastIndex + 1
			continue
		} else {
			// 最后一个最小字符前面有比它大的字符, 找到这个第一个符合这个条件的字符，和最后一个最小字符交换位置
			for i, oc := range occurence {
				if oc > baseIndex+i {
					toSwitchIndex = baseIndex + i
					break
				}
			}

			return switchChar(input, toSwitchIndex, smallestLastIndex)
		}
	}

	return input
}

func switchChar(input string, s, d int) string {
	bs := []byte(input)
	bs[s], bs[d] = bs[d], bs[s]

	return string(bs)
}
