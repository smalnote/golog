package lang

import "testing"
import "time"

type bus chan chan int

func TestChanofChan(t *testing.T) {
	b5 := make(bus, 5)
	go func() {
		for c := range b5 {
			go func() {
				for v := range c {
					print(v)
				}
			}()
		}
	}()

	go func() {
		for i := 0; i < 6; i++ {
			c := make(chan int)
			b5 <- c
			go func() {
				for j := 0; j < 10; j++ {
					c <- -1
				}
				close(c)
			}()
		}
		close(b5)
	}()
	<-time.After(5 * time.Second)
}

func TestChan(t *testing.T) {
	c := make(chan int)
	c2 := make(chan int)
	b := make(bus)

	go func() {
		b <- c
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		b <- c2
		for i := 0; i < 10; i++ {
			c2 <- i + 10
		}
		close(c2)
	}()

	go func() {
		for {
			c, ok := <-b
			if !ok {
				break
			}
			for v := range c {
				t.Log(v)
			}
		}
	}()

	<-time.After(5 * time.Second)
}
