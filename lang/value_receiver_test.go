package lang

import (
	"testing"
)

type valueReceiver struct {
	value int
}

func (v valueReceiver) Add(c int) {
	v.value += c
}

func (v valueReceiver) GetAddr() *valueReceiver {
	return &v
}

func TestValueReceiver(t *testing.T) {
	v := valueReceiver{
		value: 0,
	}
	v.Add(100)
	if v.value == 0 {
		t.Log("value receiver is a copy ")
	} else {
		t.Log("value receiver is a reference ")
	}

	if &v != v.GetAddr() {
		t.Log("calling func on a value receiver will execute the func on a copy of the receiver ")
	}
}
