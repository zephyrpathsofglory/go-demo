package apitest

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 启动 2个groutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完。
func main4() {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	c1 := make(chan struct{})
	c2 := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		go func() {
			time.Sleep(1 * time.Second)
			c1 <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			fmt.Println("task which lasts 1s timeout")
		case <-c1:
			fmt.Println("task which lasts 1s has not timeout")
		}
	}()

	go func() {
		defer wg.Done()
		go func() {
			time.Sleep(3 * time.Second)
			c2 <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			fmt.Println("task which lasts 3s timeout")
		case <-c2:
			fmt.Println("task which lasts 3s has not timeout")
		}
	}()

	wg.Wait()

	return
}
