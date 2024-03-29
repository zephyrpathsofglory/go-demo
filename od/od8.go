package od

import (
	"cmp"
	"slices"
)

/*
题目
某公司组织一场公开招聘活动，假设由于人数和场地的限制， 每人每次面试的时长不等，并已经安排给定，
用(S1,E1)、(S2,E2)、(Sj,Ej)...(Si < Ei，均为非负整数)表示每场面试的开始和结束时间。
面试采用一对一的方式，即一名面试官同时只能面试一名应试者， 一名面试官完成一次面试后可以立即进行下一场面试，且每个面试官的面试人次不超过m。
为了支撑招聘活动高效顺利进行，请你计算至少需要多少名面试官。

输入
输入的第一行为面试官的最多面试人次 m，第二行为当天总的面试场次 n， 接下来的 n 行为每场面试的起始时间和结束时间，起始时间和结束时间用空格分隔。
其中，1 <= n, m <= 500

输出
输出一个整数，表示至少需要的面试官数量。
*/

func leastInterviewer(interviews [][2]int, maxInterviewPerInterviewer, interviewCount int) int {
	schedule := make([][]int, interviewCount, interviewCount)

	slices.SortFunc(interviews, func(x, y [2]int) int {
		return cmp.Compare(x[0], y[0])
	})

	for _, interview := range interviews {
		for i, s := range schedule {
			if len(s) == 0 || (len(s) < maxInterviewPerInterviewer && s[len(s)-1] <= interview[0]) {
				schedule[i] = append(schedule[i], interview[1])
				break
			}
		}
	}

	s := slices.DeleteFunc(schedule, func(i []int) bool {
		return len(i) == 0
	})

	return len(s)
}
