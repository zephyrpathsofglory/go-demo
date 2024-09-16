/*
实现文本编辑器：a-z字母，正常输入
@ 切换大小写
+ 换行
~ 删除光标左边的字符
- 删除光标右边的字符
^ 光标上移
* 光标下移
< 光标左移
> 光标右移

默认小写
删除时需要自动合并行
换行时光标后面的自动变新行
左右移动到达行尾或者行首时，自动换行到下一行首或者上一行行尾
当在第一行上移时，移动到行首
当在最后一方下移时，移动到行尾

如果3行分别时4，2，5个字符，光标在第三行的第四个字符后面，上移光标会到第二行的最后，即第2个字符后面，
此时再往上移动光标会移动到第一行的第2个字符后面，即移动不会记录光标最初始的offset
*/
package od

import "strings"

const (
	newLineChar = "|"
)

func addString(src string, cursor int, s string) (string, int) {
	left := src[:cursor-1]
	right := src[cursor-1:]

	res := left + s + right
	cursor++

	return res, cursor
}

func deleteChar(src string, cursor int) (string, int) {
	if cursor == 1 {
		return src, cursor
	}

	left := src[0 : cursor-2]
	right := src[cursor-1:]
	res := left + right
	cursor--

	return res, cursor
}

func deleteCharRight(src string, cursor int) string {
	if cursor == len(src) {
		return src
	}

	left := src[0 : cursor-1]
	right := src[cursor:]
	res := left + right

	return res
}

func charIfToAdd(c int32, isCap bool) string {
	if isAlpha(c) {
		toAdd := string(c)
		if isCap {
			toAdd = string(toCapitalAlpha(c))
		}
		return toAdd
	}

	return ""
}

func moveCursorRight(size, cursor int) int {
	if cursor < size {
		cursor++
	}

	return cursor
}

func moveCursorLeft(cursor int) int {
	if cursor > 1 {
		cursor--
	}

	return cursor
}

func moveCursorUp(src string, cursor int) int {
	otherCursor := cursor
	offset := 0
	for cursor > 1 && string(src[cursor-2]) != newLineChar {
		otherCursor--
		cursor--
		offset++
	}

	// 在第一行，cursor 直接移动到行首
	if cursor == 1 {
		return cursor
	}

	// 跳过换行符
	cursor--
	otherCursor--

	// 双指针，先移动左边指针
	cursorNeedMove := true
	for i := 0; i < offset; i++ {
		if string(src[otherCursor-2]) != newLineChar {
			otherCursor--
		} else {
			cursorNeedMove = false
			break
		}
	}

	if !cursorNeedMove {
		return cursor
	}

	for string(src[otherCursor]) != newLineChar && otherCursor-1 != 0 {
		otherCursor--
		cursor--
	}

	return cursor
}

func moveCursorDown(src string, cursor int) int {
	offset := 0
	for i := cursor; i > 0; i-- {
		if string(src[i]) != newLineChar {
			offset++
		}
	}

	for i := cursor; i < len(src); i++ {
		cursor++
		if string(src[i]) == newLineChar {
			for j := 0; j < offset; j++ {
				if cursor < len(src) {
					cursor++
				}
			}
		}
	}

	return cursor
}

func strEdit(inputStr string) []string {
	res := ""
	cursor := 1
	isCap := false

	for _, c := range inputStr {
		toAdd := charIfToAdd(c, isCap)
		if toAdd != "" {
			res, cursor = addString(res, cursor, toAdd)
			continue
		}
		switch string(c) {
		case "+":
			res, cursor = addString(res, cursor, newLineChar)
		case "~":
			res, cursor = deleteChar(res, cursor)
		case "@":
			isCap = !isCap
		case "-":
			res = deleteCharRight(res, cursor)
		case "^":
			cursor = moveCursorUp(res, cursor)
		case "*":
			cursor = moveCursorDown(res, cursor)
		case "<":
			cursor = moveCursorLeft(cursor)
		case ">":
			cursor = moveCursorRight(len(res), cursor)
		}
	}

	return strings.Split(res, newLineChar)
}

func isAlpha(i int32) bool {
	return i >= 97 && i <= 122
}

func toCapitalAlpha(i int32) int32 {
	return i - 32
}
