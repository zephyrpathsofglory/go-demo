package trw

type maxValue struct {
	idxVal        int
	leftMaxValue  int
	rightMaxValue int
}

func TrappingRainWater(nums []int) int {
	mvs := make([]*maxValue, len(nums))

	max := 0
	for i, num := range nums {
		mv := maxValue{
			idxVal: num,
		}

		if num >= max {
			max = num
			mv.leftMaxValue = num
		} else {
			mv.leftMaxValue = max
		}

		mvs[i] = &mv
	}

	max = 0
	for i := range nums {
		mv := mvs[len(nums)-1-i]
		num := nums[len(nums)-1-i]
		if num >= max {
			max = num
			mv.rightMaxValue = num
		} else {
			mv.rightMaxValue = max
		}
	}

	total := 0
	for _, mv := range mvs {
		total += (minV(mv.leftMaxValue, mv.rightMaxValue) - mv.idxVal)
	}

	return total
}

func minV(x, y int) int {
	if x < y {
		return x
	}

	return y
}
