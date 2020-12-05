package lang

import (
	"testing"
	"time"
)

func TestRangChan(t *testing.T) {
	c := make(chan int, 4)

	go func(c chan int) {
		for v := range c {
			t.Log(v)
		}
	}(c)

	c <- 1
	close(c)
	<-time.After(5 * time.Millisecond)
}
