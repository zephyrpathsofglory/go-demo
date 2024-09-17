package apitest

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 多 producer 和 多 consumer ，通过 channel 发送，并且打印出 生产者和消费者
func Test2(_ *testing.T) {
	fmt.Println("start")
	ctx := context.TODO()
	ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	task(ctx1)
	fmt.Println("end")
}

func task(ctx context.Context) {
	var wg sync.WaitGroup
	c := make(chan string, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					c <- fmt.Sprintf("message from %d", i)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}

	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case content := <-c:
					fmt.Printf("received from %d, content: %s\n", i, content)
				}
			}
		}(j)
	}

	wg.Wait()
}
