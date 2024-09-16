package od

import (
	"strconv"
	"strings"
)

/*
分页器
pageCount	总页数
pagerCount	页码按钮的数量，当总页数超过该值时会折叠
currentPage	当前页数
*/

func pager(pageCount, pagerCount, currentPage int) string {
	if pageCount <= pagerCount {
		pages := make([]int, 0, pageCount)
		for i := 1; i <= pageCount; i++ {
			pages = append(pages, i)
		}
		return getPagerString([][]int{pages})
	}

	// pageCount > pagerCount 必然至少有一边有省略号

	hasRightDot := false
	hasLeftDot := false

	if currentPage < pageCount-pagerCount/2 {
		hasRightDot = true
	}

	if currentPage > pagerCount/2+1 {
		hasLeftDot = true
	}

	if hasLeftDot && hasRightDot {
		count := pagerCount / 2
		leftEdge := currentPage - (count - 2)
		rightEdge := leftEdge + (pagerCount - 4) - 1
		secondPart := make([]int, 0, pagerCount)
		for i := leftEdge; i <= rightEdge; i++ {
			secondPart = append(secondPart, i)
		}
		return getPagerString([][]int{
			{1},
			secondPart,
			{pageCount},
		})
	}

	if hasLeftDot && !hasRightDot {
		leftEdge := pageCount - (pagerCount - 2) + 1
		secondPart := make([]int, 0, pagerCount)
		for i := leftEdge; i <= pageCount; i++ {
			secondPart = append(secondPart, i)
		}
		return getPagerString([][]int{
			{1},
			secondPart,
		})
	}

	if !hasLeftDot && hasRightDot {
		secondPart := make([]int, 0, pagerCount)
		for i := 1; i <= pagerCount-2; i++ {
			secondPart = append(secondPart, i)
		}
		return getPagerString([][]int{
			secondPart,
			{pageCount},
		})
	}

	return ""
}

func getPagerString(pages [][]int) string {
	ss := make([]string, 0, len(pages))
	for _, p := range pages {
		str := make([]string, 0, len(p))
		for _, s := range p {
			str = append(str, strconv.Itoa(s))
		}

		ss = append(ss, strings.Join(str, " "))
	}

	return strings.Join(ss, " ... ")
}
