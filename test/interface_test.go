package test

import (
	"testing"
)

type acter interface {
	run()
	walk()
}

type person struct {
	name string
	age  int
}

func (p person) walk() {
	return
}

func (p *person) run() {
	return
}
func TestInterfaces(t *testing.T) {
	var i acter
	i = &person{}
	i.run()
}
