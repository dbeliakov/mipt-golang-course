package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func foo(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case <-time.After(time.Second):
		fmt.Println("Finished")
	case <-ctx.Done():
		fmt.Println("Aborted", ctx.Err())
	}
	wg.Done()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go foo(ctx, &wg)

	wg.Wait()
}
