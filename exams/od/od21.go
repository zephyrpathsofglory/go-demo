/*
Jessi初学英语，为了快速读出一串数字，编写程序将数字转换成英文：

具体规则如下:
1.在英语读法中三位数字看成一整体，后面再加一个计数单位。从最右边往左数，三位一单位，例如12,345 等
2.每三位数后记得带上计数单位 分别是thousand, million, billion.
3.公式：百万以下千以上的数 X thousand X, 10亿以下百万以上的数：X million X thousand X, 10 亿以上的数：
X billion X million X thousand X. 每个X分别代表三位数或两位数或一位数。
4.在英式英语中百位数和十位数之间要加and，美式英语中则会省略，我们这个题目采用加上and，百分位为零的话，这道题目我们省略and

下面再看几个数字例句：
22: twenty two
100:  one hundred
145:  one hundred and forty five
1,234:  one thousand two hundred and thirty four
8,088:  eight thousand (and) eighty eight (注:这个and可加可不加，这个题目我们选择不加)
486,669:  four hundred and eighty six thousand six hundred and sixty nine
1,652,510:  one million six hundred and fifty two thousand five hundred and ten

说明：
数字为正整数，不考虑小数，转化结果为英文小写；
保证输入的数据合法
关键字提示：and，billion，million，thousand，hundred。

数据范围：
1 <= n <= 2000000

输入描述：
输入一个long型整数

输出描述：
输出相应的英文写法

示例1
输入：
22
复制
输出：
twenty two
*/
package od

import "fmt"

var numStrMap = map[uint8]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func numberToStr(i int) string {
	billion := i / (1000 * 1000 * 1000)
	if billion >= 1000 {
		panic("too big value")
	}
	left := i % (1000 * 1000 * 1000)
	var res string
	if billion > 0 {
		res = fmt.Sprintf("%s billion", convertInnerThousandToString(uint16(billion)))
	}

	million := left / (1000 * 1000)
	left = left % (1000 * 1000)

	if million > 0 {
		millStr := fmt.Sprintf("%s million", convertInnerThousandToString(uint16(million)))
		if len(res) > 0 {
			res = fmt.Sprintf("%s %s million", res, millStr)
		} else {
			res = millStr
		}
	}

	thousand := left / 1000
	left = left % 1000

	if thousand > 0 {
		thouStr := fmt.Sprintf("%s thousand", convertInnerThousandToString(uint16(thousand)))
		if len(res) > 0 {
			res = fmt.Sprintf("%s %s", res, thouStr)
		} else {
			res = thouStr
		}
	}

	if left == 0 {
		return res
	} else {
		lastStr := convertInnerThousandToString(uint16(left))
		if len(res) == 0 {
			return lastStr
		} else {
			return fmt.Sprintf("%s %s", res, lastStr)
		}
	}
}

func convertInnerThousandToString(i uint16) string {
	var res string
	hundred := i / 100
	hundredLeft := i % 100

	if hundred > 0 {
		res = fmt.Sprintf("%s hundred", numStrMap[uint8(hundred)])
	}

	if hundredLeft == 0 {
		return res
	}

	if left, ok := numStrMap[uint8(hundredLeft)]; ok {
		if len(res) > 0 {
			res = fmt.Sprintf("%s and %s", res, left)
		} else {
			res = left
		}
	} else {
		dig := hundredLeft % 10
		tens := hundredLeft - dig
		lastTwoNumStr := fmt.Sprintf("%s %s", numStrMap[uint8(tens)], numStrMap[uint8(dig)])
		if len(res) > 0 {
			res = fmt.Sprintf("%s and %s", res, lastTwoNumStr)
		} else {
			res = lastTwoNumStr
		}
	}

	return res
}
