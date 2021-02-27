package test

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestMap(t *testing.T) {
	ma := make(map[int]person)
	andy := person{
		name: "andy",
		age:  1,
	}

	ma[1] = andy

	fmt.Println(ma[1].age)
	//ma[1].name = "andy" //Error!!

}
