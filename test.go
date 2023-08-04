package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

var n = runtime.GOMAXPROCS(1)

func main1() {
	fmt.Println(n)
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var temp uint64
	for i := 0; ; i++ {
		mb := printMemoryUsage()
		fmt.Printf("Cycle %d: %dMB\n", i, mb)
		fmt.Printf("delta: %dMB\n", mb-temp)
		temp = mb
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	fmt.Printf("buffer address: %p, size: %d, capacity: %d\n", b, b.Len(), b.Cap())
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	fmt.Printf("buffer address after grow: %p, size: %d, capacity: %d\n", b, b.Len(), b.Cap())
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

func printMemoryUsage() (mb uint64) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	fmt.Printf("Allocated memory: %d bytes\n", stats.Alloc)
	fmt.Printf("Total memory allocated and not yet freed: %d bytes\n", stats.TotalAlloc)
	fmt.Printf("Memory allocated for internal data structures: %d bytes\n", stats.Sys)
	fmt.Printf("Number of allocated objects: %d\n", stats.Mallocs)
	fmt.Printf("Number of freed objects: %d\n", stats.Frees)
	fmt.Printf("Heap memory obtained from the OS: %d bytes\n", stats.HeapAlloc)
	fmt.Printf("Heap memory obtained from the OS, but not released back: %d bytes\n", stats.HeapSys)
	fmt.Printf("Heap memory used by heap-allocated objects: %d bytes\n", stats.HeapInuse)
	fmt.Printf("Heap memory released to the OS: %d bytes\n", stats.HeapReleased)
	fmt.Printf("Heap memory idle: %d bytes\n", stats.HeapIdle)
	fmt.Printf("Heap memory reserved for future allocations: %d bytes\n", stats.HeapIdle-stats.HeapReleased)
	fmt.Printf("Stack memory in use: %d bytes\n", stats.StackInuse)
	fmt.Printf("Stack memory obtained from the OS: %d bytes\n", stats.StackSys)
	fmt.Printf("MSpan memory obtained from the OS: %d bytes\n", stats.MSpanSys)
	fmt.Printf("MCache memory obtained from the OS: %d bytes\n", stats.MCacheSys)
	fmt.Printf("GC memory obtained from the OS: %d bytes\n", stats.GCSys)
	fmt.Printf("Number of completed garbage collection cycles: %d\n", stats.NumGC)
	fmt.Printf("Total time spent in garbage collection: %v\n", time.Duration(stats.PauseTotalNs))

	return stats.Alloc / 1024 / 1024
}
