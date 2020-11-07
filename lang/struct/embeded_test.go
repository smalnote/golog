package structtest

import (
	"fmt"
	"testing"
)

type Collection struct {
	size int
}

func (c Collection) Size() int {
	return c.size
}

type List struct {
	Collection
}

type ArrayList struct {
	List
	Values []int
}

type ICollection interface {
	Size() int
}

func PrintSize(c ICollection) {
	fmt.Printf("ICollection's size = %d\n", c.Size())
}

func TestEmbeded(t *testing.T) {
	a := ArrayList{
		List: List{
			Collection: Collection{
				100,
			},
		},
		Values: nil,
	}
	fmt.Printf("a's type is %T\n", a)
	fmt.Printf("a's value is %v\n", a)
	fmt.Printf("a.size = %d\n", a.size)
	fmt.Printf("a.List.size = %d\n", a.List.size)
	fmt.Printf("a.List.Collection.size = %d\n", a.List.Collection.size)
	fmt.Printf("a.Size() = %d\n", a.Size())
	PrintSize(a) // a's embeded List's embeded's Collection has method Size(), so a is a ICollection
}
