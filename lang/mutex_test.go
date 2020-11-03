package lang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mutex sync.RWMutex
var count int

func TestRWMutex(t *testing.T) {
	go A()
	time.Sleep(2 * time.Second)
	mutex.Lock()
	defer mutex.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	mutex.RLock()
	defer mutex.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mutex.RLock()
	defer mutex.RUnlock()
}
