package lang

import "testing"

func TestTypeAssertion(t *testing.T) {
	values := []interface{}{
		1,
		"a",
		5.1,
		true,
		'b',
		make(chan int),
		func() {},
		make([]int, 10),
	}

	for _, v := range values {
		switch v.(type) {
		case int:
			vv, _ := v.(int)
			t.Logf("v is a int, %d\n", vv)
		case string:
			vv, _ := v.(string)
			t.Logf("v is a string, %s\n", vv)
		case float64:
			vv, _ := v.(float64)
			t.Logf("v is a float64, %f\n", vv)
		case byte:
			vv, _ := v.(byte)
			t.Logf("v is a byte, %d\n", vv)
		case bool:
			vv, _ := v.(bool)
			t.Logf("v is a bool, %t\n", vv)
		case chan int:
			vv, _ := v.(chan int)
			t.Logf("v is a chan int, %v\n", vv)
		case func():
			vv, _ := v.(func())
			t.Logf("v is a func(), %v\n", &vv)
		case []int:
			vv, _ := v.([]int)
			t.Logf("v is a slice int, %v\n", vv)
		}
	}
}
