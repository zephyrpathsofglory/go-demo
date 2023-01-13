package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	round := gcd(n, k)

	for i := 0; i < round; i++ {
		cur := i
		tmp := nums[cur]
		for nxt := (i + k) % n; nxt != i; nxt = (nxt + k) % n {
			t := nums[nxt]
			nums[nxt] = tmp
			tmp = t
		}
		nums[i] = tmp
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
