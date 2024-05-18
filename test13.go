package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex 可以被其他 goroutine Unlock
func main() {
	var m sync.Mutex

	go func() {
		time.Sleep(5 * time.Second)
		m.Unlock()
		fmt.Println("unlocked")
	}()

	m.Lock()

	time.Sleep(10 * time.Second)
	m.Unlock()
}
