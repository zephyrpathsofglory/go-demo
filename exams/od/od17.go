/*
给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。
对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，
则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换。
输入描述:
输入为两行，第一行为参数K，第二行为字符串S。
输出描述:
输出转换后的字符串。
示例1
输入
3
12abc-abCABc-4aB@
输出
12abc-abc-ABC-4aB-@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每3个字符一组为abC、ABc、4aB、@，abC中小写字母较多，转换为abc，ABc中大写字母较多，
转换为ABC，4aB中大小写字母都为1个，不做转换，@中没有字母，连起来即12abc-abc-ABC-4aB-@
示例2
输入
12
12abc-abCABc-4aB@
输出
12abc-abCABc4aB@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每12个字符一组为abCABc4aB@，这个子串中大小写字母都为4个，不做转换，连起来即12abc-abCABc4aB@
*/

package od

import "strings"

func SplitString(s string, n int) string {
	cur := 0
	splitChar := "-"
	for i, c := range s {
		if string(c) != splitChar {
			continue
		}

		cur = i
		break
	}

	res := []string{s[:cur]}

	counter := 0
	buffer := ""
	letterCount := 0
	capitalLetterCount := 0
	for i := cur + 1; i < len(s); i++ {
		if counter == n {
			if letterCount > capitalLetterCount {
				buffer = iterateToLetter(buffer, false)
			} else if letterCount < capitalLetterCount {
				buffer = iterateToLetter(buffer, true)
			}

			res = append(res, buffer)

			buffer = ""
			capitalLetterCount = 0
			letterCount = 0
			counter = 0
		}

		c := string(s[i])
		if c == splitChar {
			continue
		}

		buffer += c

		if isLetter(int32(s[i])) {
			letterCount++
		} else if isCapitalLetter(int32(s[i])) {
			capitalLetterCount++
		}

		counter++
	}

	if len(buffer) != 0 {
		res = append(res, buffer)
	}

	return strings.Join(res, splitChar)
}

func iterateToLetter(input string, isCap bool) string {
	for idx, i := range input {
		if isCap {
			input = input[0:idx] + toCapitalLetter(i) + input[idx+1:]
		} else {
			input = input[0:idx] + toLetter(i) + input[idx+1:]
		}
	}

	return input
}

func toCapitalLetter(i int32) string {
	if isCapitalLetter(i) {
		return string(i)
	} else {
		return string(i - 32)
	}
}

func isLetter(i int32) bool {
	return i >= 97 && i <= 122
}

func isCapitalLetter(i int32) bool {
	return i >= 65 && i <= 90
}

func toLetter(i int32) string {
	if isLetter(i) {
		return string(i)
	} else {
		return string(i + 32)
	}
}
