package main

import "fmt"

// 当select监控多个chan同时到达就绪态时，如何先执行某个任务？
func main6() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})

	select {
	case <-c1:
		fmt.Print("c1")
	case <-c2:
		for {
			select {
			case <-c1:
				fmt.Print("c1, c2 received at the same, use c1")
			default:
				fmt.Print("c2")
				return
			}
		}
	}
}
