package main

import "fmt"

func MyAppend(arr []int) []int {
	arr = append(arr, 5)

	return arr
}

func MyAdd(arr []int) {
	for i := range arr {
		arr[i] = arr[i] + 5
	}

	return
}

func main9() {
	arr := []int{1, 2, 3, 4}
	MyAppend(arr)
	fmt.Println(arr)

	MyAdd(arr)
	fmt.Println(arr)
}
