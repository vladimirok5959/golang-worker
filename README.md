# golang-worker
Simple background worker

## How to use
```
go get github.com/vladimirok5959/golang-worker
```
```
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vladimirok5959/golang-worker/worker"
)

func main() {
	fmt.Printf("Start!\n")

	w1 := worker.New(func(ctx context.Context, w *worker.Worker, o *[]worker.Iface) {
		fmt.Printf("Worker #1 one tick\n")
		for i := 0; i < 1000; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker #1 fine I will shutdown!\n")
				return
			default:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}, nil)

	w2 := worker.New(func(ctx context.Context, w *worker.Worker, o *[]worker.Iface) {
		fmt.Printf("Worker #2 one tick\n")
		for i := 0; i < 1000; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker #2 fine I will shutdown!\n")
				return
			default:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}, nil)

	_ = worker.New(func(ctx context.Context, w *worker.Worker, o *[]worker.Iface) {
		fmt.Printf("Worker #3 one tick\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker #3 Exit\n")
		w.Shutdown(nil)
	}, nil)

	// Just wait for goroutines
	time.Sleep(3 * time.Second)

	// Shutdown first
	// Normally, it's must be used with context
	w1.Shutdown(nil)

	// Shutdown second
	// Normally, it's must be used with context
	if err := w2.Shutdown(nil); err != nil {
		fmt.Printf("Worker #2 shutdown error: %s\n", err.Error())
	}

	// Wait for third
	// Will be exited automatically
	time.Sleep(1 * time.Second)

	fmt.Printf("End!\n")
}
```
