package lang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mu sync.RWMutex
var count int

func TestRWMutex(t *testing.T) {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
}
