// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	} else if n == 0 {
		return 1
	} else {
		k := n % 2
		hf := n / 2
		if k == 0 {
			r := myPow(x, hf)
			return r * r
		} else {
			r := myPow(x, hf)
			return r * r * x
		}
	}
}
