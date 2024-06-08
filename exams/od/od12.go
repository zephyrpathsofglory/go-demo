package od

/*
为了提升数据传输的效率，会对传输的报文进行压缩处理。输入一个压缩后的报文，请返回它解压后的原始报文。
压缩规则：n[str]，表示方括号内部的 str 正好重复 n 次。注意 n 为正整数（0 < n <= 100），str只包含小写英文字母，不考虑异常情况。
" “输入描述:
输入压缩后的报文：
1）不考虑无效的输入，报文没有额外的空格，方括号总是符合格式要求的；
2）原始报文不包含数字，所有的数字只表示重复的次数 n ，例如不会出现像 5b 或 3[8] 的输入；
输出描述:
解压后的原始报文
注：
1）原始报文长度不会超过1000，不考虑异常的情况
---------------------------
示例1
输入
3[k]2[mn]
输出
kkkmnmn
说明
k 重复3次，mn 重复2次，最终得到 kkkmnmn
---------------------------
示例2
输入
3[m2[c]]
输出
mccmccmcc
说明
m2[c] 解压缩后为 mcc，重复三次为 mccmccmcc”
*/

func decompression(compressed string) string {
	type appearance struct {
		body      string
		times     int
		multibody string
	}

	stack := []*appearance{}

	compressed = "1[" + compressed + "]"

	for _, c := range compressed {
		if n, ok := IsNumber(c); ok {
			stack = append(stack, &appearance{
				times: n,
			})
			continue
		}

		if IsLeftBracket(c) {
			continue
		}

		lastStack := stack[len(stack)-1]
		if IsLetter(c) {
			if lastStack.times > 0 {
				lastStack.multibody += string(c)
			} else {
				lastStack.body += string(c)
			}

			continue
		}

		if IsRightBracket(c) {
			for i := 0; i < lastStack.times; i++ {
				lastStack.body += string(lastStack.multibody)
			}

			if len(stack) == 1 {
				return stack[0].body
			}

			previousStack := stack[len(stack)-2]
			previousStack.multibody += lastStack.body

			stack = stack[0 : len(stack)-1]

			continue
		}
	}

	return stack[0].body
}

func IsNumber(c rune) (int, bool) {
	var zero rune = '0'

	if c-zero > 0 && c-zero < 10 {
		return int(c - zero), true
	}

	return 0, false
}

func IsLetter(c rune) bool {
	var a rune = 'a'
	var z rune = 'z'
	return c >= a && c <= z
}

func IsLeftBracket(c rune) bool {
	var l rune = '['

	return c == l
}

func IsRightBracket(c rune) bool {
	var r rune = ']'

	return c == r
}
