package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 请使用Go语言实现sync.WaitGroup的三个功能：Add、Done、Wait

type CustomWaitGroup struct {
	count int32
	done  chan struct{}
}

func (cwg CustomWaitGroup) Add(i int32) {
	atomic.AddInt32(&cwg.count, i)
	if atomic.LoadInt32(&cwg.count) < 0 {
		panic("negative counter")
	}
}

func (cwg CustomWaitGroup) Done() {
	cwg.Add(-1)
}

func (cwg CustomWaitGroup) Wait() {
	<-cwg.done
}

func main7() {
	var wg CustomWaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(2)
		wg.Done()
	}()

	fmt.Print("start")
	wg.Wait()
	return
}
