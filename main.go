package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vladimirok5959/golang-worker/worker"
)

func main() {
	fmt.Printf("Start!\n")

	w1 := worker.New(func(ctx context.Context, w *worker.Worker) {
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
	})

	w2 := worker.New(func(ctx context.Context, w *worker.Worker) {
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
	})

	_ = worker.New(func(ctx context.Context, w *worker.Worker) {
		fmt.Printf("Worker #3 one tick\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker #3 Exit\n")
		w.Shutdown(nil)
	})

	time.Sleep(3 * time.Second)

	w1.Shutdown(nil)

	if err := w2.Shutdown(nil); err != nil {
		fmt.Printf("Worker #2 shutdown error: %s\n", err.Error())
	}

	time.Sleep(1 * time.Second)

	fmt.Printf("End!\n")
}
