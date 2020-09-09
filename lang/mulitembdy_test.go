package lang

import "testing"

type Greeting interface {
	SayHi() string
}

type Duck struct{}

func (d *Duck) SayHi() string {
	return "quack"
}

type Fish struct{}

func (f *Fish) SayHi() string {
	return "blub"
}

type DuckFish struct {
	Duck
	f Fish
}

func TestMultiEmbodded(t *testing.T) {
	var d Greeting
	d = &DuckFish{}
	println(d.SayHi())
}
