package test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	b := &a
	fmt.Println(b[0])

	// c := []int{1, 2, 3, 4}
	// d := &c
	// fmt.Println(d[0])
}
