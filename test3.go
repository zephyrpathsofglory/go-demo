package main

import (
	"fmt"
	"sync"
)

// 有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。

func main3() {
	var wg sync.WaitGroup
	cat := make(chan struct{})
	dog := make(chan struct{})
	fish := make(chan struct{})
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-cat
			fmt.Println("cat")
			dog <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-dog
			fmt.Println("dog")
			fish <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-fish
			fmt.Println("fish")
			if i == 99 {
				return
			}
			cat <- struct{}{}
		}
	}()

	cat <- struct{}{}

	wg.Wait()

	return
}
