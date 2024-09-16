package apitest

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"syscall"
	"testing"
)

// Define functions fn1 through fn5, each performing simple operations // and printing results.

func fn1(wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion of goroutine
	fmt.Println("Executing function 1")
}

func fn2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing function 2")
}

func fn3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing function 3")
}

func fn4(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing function 4")
}

func fn5(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing function 5")
}

func fn6(wg *sync.WaitGroup) {
	defer wg.Done()
	val, found := syscall.Getenv("HOME")
	fmt.Println(val, found)
	fmt.Println("Executing function 6")
}

func fn7(wg *sync.WaitGroup) {
	defer wg.Done()
	content, err := os.ReadFile("./data.txt")
	fmt.Println(string(content), err)
	fmt.Println("Executing function 5")
}

// Define arithmetic operations: sum, subtract, multiply, divide

func Test14(_ *testing.T) {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(3) // Limit the number of OS threads

	// Create and open a file to store trace data
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start tracing
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// Adding 5 goroutines to the WaitGroup
	wg.Add(7)

	// Starting each function as a goroutine
	go fn1(&wg)
	go fn2(&wg)
	go fn3(&wg)
	go fn4(&wg)
	go fn5(&wg)
	go fn6(&wg)
	go fn7(&wg)

	// Waiting for all goroutines to complete
	wg.Wait()

	// Stop tracing
	fmt.Println("All goroutines completed")
}
