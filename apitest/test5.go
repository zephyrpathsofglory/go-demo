package apitest

import (
	"fmt"
	"sync"
)

// 两个协程交替打印10个字母和数字
func main5() {
	numChan := make(chan struct{})
	charChan := make(chan struct{})
	times := 10

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			<-charChan
			fmt.Println("a")
			numChan <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			<-numChan
			fmt.Println(1)
			if i == times-1 {
				return
			}
			charChan <- struct{}{}
		}
	}()

	charChan <- struct{}{}
	wg.Wait()

	return
}
