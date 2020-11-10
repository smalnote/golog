package lang

import (
	"context"
	"testing"
	"time"
)

type contextKey1 struct{}
type contextKey2 struct{}

func TestValue(t *testing.T) {
	var ctx context.Context

	ctx = context.Background()

	key1 := contextKey1{}
	key2 := contextKey2{}

	t.Log(ctx.Value(key1))
	ctx = context.WithValue(ctx, key1, "AA")
	t.Log(ctx.Value(key1))
	ctx = context.WithValue(ctx, key2, "AAA")
	t.Log(ctx.Value(key2))
}

func TestTimeout(t *testing.T) {
	var ctx context.Context

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(51 * time.Millisecond):
				t.Log("running... ")
			case <-ctx.Done():
				t.Log("canceled! ")
				return
			}

		}
	}(ctx)

	<-ctx.Done()

}

func TestCancel(t *testing.T) {
	ctx := context.Background()

	ctx, c1 := context.WithCancel(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(10 * time.Millisecond):
				t.Log("hello ")
			case <-ctx.Done():
				t.Log("hello canceled ")
				return
			}
		}
	}(ctx)

	ctx, c2 := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(30 * time.Millisecond):
				t.Log("hi ")
			case <-ctx.Done():
				t.Log("hi canceled ")
				return
			}
		}
	}(ctx)

	<-time.After(100 * time.Millisecond)
	c1()
	<-time.After(100 * time.Millisecond)
	c2()

}
