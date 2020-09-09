package lang

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func TestIota(t *testing.T) {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	var v interface{} = mutexLocked
	r := reflect.ValueOf(v)
	fmt.Println(r)
	fmt.Println(r.Kind())
	fmt.Println(2 & mutexStarving)
}
