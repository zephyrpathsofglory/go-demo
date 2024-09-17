package apitest

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// sync.Mutex 可以被其他 goroutine Unlock
func Test13(_ *testing.T) {
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

func TestSliceIdx(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[3:]

	assert.Len(t, b, 0)
}
