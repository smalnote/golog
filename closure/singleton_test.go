package closure

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	c := make(chan *int)

	go getSingleton(c)
	go getSingleton(c)

	v1 := <-c
	v2 := <-c

	t.Logf("v1 = %v, v2 = %v, equals %t\n", v1, v2, v1 == v2)

}

func getSingleton(c chan *int) {
	c <- Singleton()
}
