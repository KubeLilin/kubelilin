package tests

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	slice := make([]int, 100)
	p := &slice

	slice[0] = 111
	slice[1] = 222

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println((*p)[0])
	fmt.Println((*p)[1])

}

func TestPointer(t *testing.T) {
	name := "deployment"
	var ss *string
	ss = &name

	println(*ss)

}
