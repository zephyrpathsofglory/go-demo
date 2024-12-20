/*
描述
IP地址是由4个0-255之间的整数构成的，用"."符号相连。
二进制的IP地址格式有32位，例如：10000011，01101011，00000011，00011000;每八位用十进制表示就是131.107.3.24
子网掩码是用来判断任意两台计算机的IP地址是否属于同一子网络的根据。
子网掩码与IP地址结构相同，是32位二进制数，由1和0组成，且1和0分别连续，其中网络号部分全为“1”和主机号部分全为“0”。
你可以简单的认为子网掩码是一串连续的1和一串连续的0拼接而成的32位二进制数，左边部分都是1，右边部分都是0。
利用子网掩码可以判断两台主机是否在同一子网中。
若两台主机的IP地址分别与它们的子网掩码进行逻辑“与”运算（按位与/AND）后的结果相同，则说明这两台主机在同一子网中。

示例：
I P 地址　 192.168.0.1
子网掩码　 255.255.255.0

转化为二进制进行运算：

I P 地址　  11000000.10101000.00000000.00000001
子网掩码　11111111.11111111.11111111.00000000

AND运算   11000000.10101000.00000000.00000000

转化为十进制后为：
192.168.0.0


I P 地址　 192.168.0.254
子网掩码　 255.255.255.0


转化为二进制进行运算：

I P 地址　11000000.10101000.00000000.11111110
子网掩码  11111111.11111111.11111111.00000000

AND运算  11000000.10101000.00000000.00000000

转化为十进制后为：
192.168.0.0

通过以上对两台计算机IP地址与子网掩码的AND运算后，我们可以看到它运算结果是一样的。均为192.168.0.0，所以这二台计算机可视为是同一子网络。

输入一个子网掩码以及两个ip地址，判断这两个ip地址是否是一个子网络。
若IP地址或子网掩码格式非法则输出1，若IP1与IP2属于同一子网络输出0，若IP1与IP2不属于同一子网络输出2。

注:
有效掩码与IP的性质为：
1. 掩码与IP每一段在 0 - 255 之间
2. 掩码的二进制字符串前缀为网络号，都由‘1’组成；后缀为主机号，都由'0'组成

输入描述：
3行输入，第1行是输入子网掩码、第2，3行是输入两个ip地址
题目的示例中给出了三组数据，但是在实际提交时，你的程序可以只处理一组数据（3行）。

输出描述：
若IP地址或子网掩码格式非法则输出1，若IP1与IP2属于同一子网络输出0，若IP1与IP2不属于同一子网络输出2

示例1
输入：
255.255.255.0
192.168.224.256
192.168.10.4
255.0.0.0
193.194.202.15
232.43.7.59
255.255.255.0
192.168.0.254
192.168.0.1
输出：
1
2
0

说明：
对于第一个例子:
255.255.255.0
192.168.224.256
192.168.10.4
其中IP:192.168.224.256不合法，输出1

对于第二个例子:
255.0.0.0
193.194.202.15
232.43.7.59
2个与运算之后，不在同一个子网，输出2

对于第三个例子，2个与运算之后，如题目描述所示，在同一个子网，输出0
*/

package od

import (
	"strconv"
	"strings"
)

func inSameSubNet(ip1, ip2, mask string) uint8 {
	ok, ip1Parts := isValidIP(ip1)
	if !ok {
		return 1
	}

	ok, ip2Parts := isValidIP(ip2)
	if !ok {
		return 1
	}

	ok, maskParts := isValidMask(mask)
	if !ok {
		return 1
	}

	for i := 0; i < len(maskParts); i++ {
		if !partMatch(ip1Parts[i], ip2Parts[i], maskParts[i]) {
			return 2
		}
	}

	return 0
}

func partMatch(i1, i2, mask uint8) bool {
	return (i1 & mask) == (i2 & mask)
}

func isValidIP(ip string) (bool, []uint8) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false, []uint8{}
	}

	numParts := make([]uint8, 0, len(parts))
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err != nil {
			return false, []uint8{}
		}

		if i < 0 || i > 255 {
			return false, []uint8{}
		}

		numParts = append(numParts, uint8(i))
	}
	return true, numParts
}

func isValidMask(mask string) (bool, []uint8) {
	ok, parts := isValidIP(mask)
	if !ok {
		return ok, parts
	}

	var res int

	for i := 0; i < 4; i++ {
		bits := 8 * (3 - i)
		operator := int(parts[i])
		k := (operator << bits)
		res += k
	}

	ok = isValidMaskVal(res)
	if ok {
		return true, parts
	} else {
		return false, []uint8{}
	}
}

func isValidMaskVal(t int) bool {
	for i := 31; i >= 0; i-- {
		t = t - 1<<i
		if t == 0 {
			break
		}

		if t < 0 {
			return false
		}
	}

	return true
}
