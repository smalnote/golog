package pipeline

import (
	"context"
	"log"
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	in := gen(ctx, 100)

	w1 := work(ctx, in)
	w2 := work(ctx, in)
	w3 := work(ctx, in)

	out := merge(ctx, w1, w2, w3)

	// 消费的数量小于, 等于, 大于生产数量都可能正常结束
	consumeN := 100
	for i := 0; i < consumeN; i++ {
		v, ok := <-out
		if ok {
			t.Logf("out: %v", v)
		} else {
			break
		}
	}

	cancel()
	<-out
}

func merge(ctx context.Context, ins ...<-chan int) <-chan int {
	wg := sync.WaitGroup{}
	out := make(chan int)

	output := func(ctx context.Context, in <-chan int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if ok {
					out <- v
				} else {
					return
				}
			}
		}
	}

	wg.Add(len(ins))
	for _, in := range ins {
		go output(ctx, in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func work(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func(in <-chan int) {
		defer log.Println("worker completed")
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if ok {
					out <- v * v
				} else {
					return
				}
			}
		}

	}(in)
	return out
}

func gen(ctx context.Context, count int) <-chan int {
	c := make(chan int)
	go func(count int) {
		defer close(c)
		for i := 1; i <= count; i++ {
			select {
			case <-ctx.Done():
				return
			case c <- i:
			}
		}
	}(count)
	return c
}
