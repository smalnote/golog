package resourcepool

import (
	"context"
	"sync"
	"testing"
	"time"
)

type conn struct {
	name string
}

func (c conn) Close() {}

func TestConnPool(t *testing.T) {
	pool := New()

	pool.Put(&conn{"a"})
	pool.Put(&conn{"b"})
	pool.Put(&conn{"c"})

	time.AfterFunc(100*time.Millisecond, func() { pool.Close() })

	wg := sync.WaitGroup{}

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go task(t, &wg, pool)
	}

	wg.Wait()
}

func task(t *testing.T, wg *sync.WaitGroup, pool ResourcePool) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 140*time.Millisecond)
	defer cancel()
	c, err := pool.GetResource(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	time.AfterFunc(100*time.Millisecond, func() {
		pool.Put(c)
	})
	t.Logf("got conn: %+v", c.(*conn))
}
