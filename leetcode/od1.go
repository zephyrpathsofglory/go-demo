package leetcode

/*
定义：开头和结尾都是元音字母（aeiouAEIOU）的字符串为元音字符串，其中混杂的非元音字母数量为瑕疵度。比如：

“a”、“aa”是元音字符串，其瑕疵度都为0

“aiur”不是元音字符串（结尾不是元音字符）

“abira”是元音字符串，其瑕疵度为2

给定一个字符串，请找出指定瑕疵度的最长元音字符子串，并输出其长度，如果找不到满足条件的元音字符子串，输出0.

例如： 输入： 1 aabeebuu, 输出： 5 （aabee， eebuu）
*/
func findSubString(input string, flawDegree int) (maxLength int) {
	started := false
	frontPointer := 0
	nextFrontPointer := 0
	nextFrontPointerSet := false
	degree := 0

	for i, char := range input {
		if isVowel(string(char)) {
			if started {
				if degree > flawDegree {
					frontPointer = nextFrontPointer
					nextFrontPointerSet = false
					degree = degree - flawDegree
				} else if degree == flawDegree {
					if !nextFrontPointerSet {
						nextFrontPointerSet = true
						nextFrontPointer = i
					}
					length := i - frontPointer + 1
					if maxLength < length {
						maxLength = length
					}
				}
			} else {
				started = true
				frontPointer = i
			}
		} else {
			if started {
				degree += 1
			} else {
				continue
			}
		}
	}

	return
}

func isVowel(char string) bool {
	vowels := []string{
		"a", "A", "e", "E", "i", "I", "o", "O", "u", "U",
	}

	for _, c := range vowels {
		if c == char {
			return true
		}
	}

	return false
}
